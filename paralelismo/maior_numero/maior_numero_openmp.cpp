#include <iostream>
#include <stdio.h>
#include <omp.h>
#include <math.h>
#include <vector>
#include <cstdlib>
#include <ctime>
#include <chrono>

int main ()
{
    // Tamanho N, simulando input do usuário.
    int N = 1200000000;

    // Análise de quantas threads estão disponíveis.
    int nthreads = omp_get_max_threads();
    std::cout << "Total de threads disponíveis: " << nthreads << std::endl << std::endl;

    // Criação do vetor de tamanho N.
    std::vector<int> vec(N);

    // Populando o vetor.
    std::srand(std::time(0));
    for (int i = 0; i < N; ++i) {
        vec[i] = std::rand() % 901; // número aleatório entre 0 e 900
    }
    vec[20] = 999; // Maior número sendo colocado manualmente em uma posição.

    // Tamanho do subvetor analisado.
    int subvec_size = vec.size() / nthreads;

    // Respostas encontradas por cada thread.
    std::vector<int> vec_respostas(nthreads);

    auto t1 = std::chrono::high_resolution_clock::now();

    #pragma omp parallel
    {
        // 0 .. nthreads-1
        int thread_id = omp_get_thread_num();

        // Seleção do subvetor analisado pela thread.
        int inicio      = thread_id*subvec_size, 
            fim         = thread_id*subvec_size+subvec_size-1,
            maior_atual = -1;

        for (size_t i = inicio; i < fim; i++)
        {
            if (vec[i] > maior_atual) {
                maior_atual = vec[i];
            }
        }

        printf("thread id: %d\nintervalo analisado: [%d..%d]\nmaior encontrado: %d\n\n", thread_id, inicio, fim, maior_atual);
        vec_respostas[thread_id] = maior_atual;
    }

    int maior_de_todos = -1;

    for (size_t i = 0; i < vec_respostas.size(); i++)
    {
        if (vec_respostas[i] > maior_de_todos) {
            maior_de_todos = vec_respostas[i];
        }
    }

    auto t2 = std::chrono::high_resolution_clock::now();
    auto duracao = std::chrono::duration_cast<std::chrono::microseconds> (t2 - t1).count();

    std::cout << "maior número: " << maior_de_todos << std::endl;
    std::cout << "tempo: " << duracao / 1000.0 << " ms" << std::endl;

    return 0;
}
