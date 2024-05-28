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

    // Criação do vetor de tamanho N.
    std::vector<int> vec(N);

    // Populando o vetor.
    std::srand(std::time(0));
    for (int i = 0; i < N; ++i) {
        vec[i] = std::rand() % 901; // número aleatório entre 0 e 900
    }
    vec[20] = 999; // Maior número sendo colocado manualmente em uma posição.

    auto t1 = std::chrono::high_resolution_clock::now();

    int maior_de_todos = -1;

    for (size_t i = 0; i < vec.size(); i++)
    {
        if (vec[i] > maior_de_todos) {
            maior_de_todos = vec[i];
        }
    }

    auto t2 = std::chrono::high_resolution_clock::now();
    auto duracao = std::chrono::duration_cast<std::chrono::microseconds>(t2 - t1).count();

    std::cout << "maior número: " << maior_de_todos << std::endl;
    std::cout << "tempo: " << duracao / 1000.0 << " ms" << std::endl;

    return 0;
}
