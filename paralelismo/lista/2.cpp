#include <mpi.h>
#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>
void crivoDeEratostenes(int n, int rank, int size)
{
    double inicio_tempo_total, fim_tempo_total;
    // Sincronizar todos os processos
    MPI_Barrier(MPI_COMM_WORLD); 
    
    inicio_tempo_total = MPI_Wtime();

    int raiz_n = static_cast<int>(std::sqrt(n));
    std::vector<char> primo(n + 1, 1);
    primo[0] = primo[1] = 0;
    
    // Primeiro processo faz o crivo até sqrt(n)
    if (rank == 0)
    {
        for (int p = 2; p <= raiz_n; ++p)
        {
            if (primo[p])
            {
                for (int i = p * p; i <= n; i += p)
                {
                    primo[i] = 0;
                }
            }
        }
    }

    // Broadcast do vetor até sqrt(n) para todos os processos
    MPI_Bcast(&primo[0], raiz_n + 1, MPI_CHAR, 0, MPI_COMM_WORLD);

    // Cada processo trabalha na parte do vetor que lhe foi designada
    int inicio = rank * (n - raiz_n) / size + raiz_n + 1;
    int fim = (rank + 1) * (n - raiz_n) / size + raiz_n;
    if (rank == size - 1)
    {
        fim = n;
    }

    for (int p = 2; p <= raiz_n; ++p)
    {
        if (primo[p])
        {
            int menor_mult = std::max(p * p, (inicio + p - 1) / p *
                                                 p);
            for (int i = menor_mult; i <= fim; i += p)
            {
                primo[i] = 0;
            }
        }
    }

    // Reunir os resultados no processo 0
    if (rank != 0)
    {
        MPI_Send(&primo[inicio], fim - inicio + 1, MPI_CHAR, 0, 0,
                 MPI_COMM_WORLD);
    }
    else
    {
        for (int p = 1; p < size; ++p)
        {
            int inicio_p = p * (n - raiz_n) / size + raiz_n + 1;
            int fim_p = (p + 1) * (n - raiz_n) / size + raiz_n;
            if (p == size - 1)
            {
                fim_p = n;
            }
            MPI_Recv(&primo[inicio_p], fim_p - inicio_p + 1,
                     MPI_CHAR, p, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        }
    }

    fim_tempo_total = MPI_Wtime();

    // Processo 0 imprime os resultados e tempo de execução
    if (rank == 0)
    {
        std::cout << "Numeros primos ate " << n << ":\n";
        for (int i = 2; i <= n; ++i)
        {
            if (primo[i])
            {
                std::cout << i << " ";
            }
        }
        std::cout << std::endl;
        std::cout << "Tempo total de execucao: " << (fim_tempo_total - inicio_tempo_total) << " segundos" << std::endl;
    }
}

int main(int argc, char *argv[])
{
    MPI_Init(&argc, &argv);

    int rank, size;
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    MPI_Comm_size(MPI_COMM_WORLD, &size);

    int n = 1000000;

    crivoDeEratostenes(n, rank, size);

    MPI_Finalize();
    
    return 0;
}