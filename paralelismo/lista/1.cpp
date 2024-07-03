#include <mpi.h>
#include <iostream>
#include <vector>
#include <algorithm>
#include <climits>

void encontrarDoisMaiores(const std::vector<int> &numeros, int &maior, int &segundoMaior)
{
    maior = segundoMaior = INT_MIN;
    for (int numero : numeros)
    {
        if (numero > maior)
        {
            segundoMaior = maior;
            maior = numero;
        }
        else if (numero > segundoMaior)
        {
            segundoMaior = numero;
        }
    }
}

int main(int argc, char **argv)
{
    MPI_Init(&argc, &argv);

    int world_size;
    MPI_Comm_size(MPI_COMM_WORLD, &world_size);

    int world_rank;
    MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);

    const int n = 100000000;

    std::vector<int> numeros;
    
    if (world_rank == 0)
    {
        // Processo mestre gera o conjunto de números
        numeros.resize(n);
        for (int i = 0; i < n; ++i)
        {
            numeros[i] = rand() % 100000000; // Números aleatórios entre 0 e ?
        }
    }

    int local_n = n / world_size;
    std::vector<int> numeros_locais(local_n);

    // Iniciar a medição de tempo
    double tempo_inicial = MPI_Wtime();

    // Distribuir os números para todos os processos
    MPI_Scatter(numeros.data(), local_n, MPI_INT, numeros_locais.data(), local_n, MPI_INT, 0, MPI_COMM_WORLD);

    // Encontrar os dois maiores números localmente
    int maior_local, segundoMaior_local;
    encontrarDoisMaiores(numeros_locais, maior_local, segundoMaior_local);

    // Reunir os resultados locais nos arrays
    std::vector<int> maiores_locais(world_size);
    std::vector<int> segundosMaiores_locais(world_size);
    
    MPI_Gather(&maior_local, 1, MPI_INT, maiores_locais.data(), 1, MPI_INT, 0, MPI_COMM_WORLD);
    MPI_Gather(&segundoMaior_local, 1, MPI_INT, segundosMaiores_locais.data(), 1, MPI_INT, 0, MPI_COMM_WORLD);

    if (world_rank == 0)
    {
        // Encontrar os dois maiores números entre os resultados locais.
        int maior_global, segundoMaior_global;

        encontrarDoisMaiores(maiores_locais, maior_global, segundoMaior_global);
        encontrarDoisMaiores(segundosMaiores_locais, maior_local, segundoMaior_local);

        if (maior_local > segundoMaior_global)
        {
            segundoMaior_global = maior_local;
        }

        // Parar a medição de tempo.
        double tempo_final = MPI_Wtime();
        double tempo_execucao = tempo_final - tempo_inicial;

        std::cout << "Os dois maiores numeros sao: " << maior_global << " e " << segundoMaior_global << std::endl;
        std::cout << "Tempo de execucao: " << tempo_execucao << "segundos " << std::endl;
    }

    MPI_Finalize();
    
    return 0;
}