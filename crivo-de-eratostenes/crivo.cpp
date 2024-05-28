#include <iostream>
#include <math.h>
#include <vector>
#include <chrono>

using namespace std;

int main() {
    // Espaço analisado.
    int valorLimite = 100000;

    // Raíz quadrada do valor limite analisado, determinando o maior número a ser analisado.
    int multiploLimite = floor(sqrt(valorLimite));

    // Considero que todos os números são primos, inicialmente.
    vector<bool> isPrime(valorLimite, true);

    // Removo os valores inicias (0 e 1), que sei que não são primos.
    isPrime[0] = false;
    isPrime[1] = false;

    auto tempoInicial = chrono::high_resolution_clock::now();

    // Começo a rodar o algoritmo do crivo à partir do primeiro primo conhecido (2).
    for (int valorAnalisado = 2; valorAnalisado < multiploLimite; valorAnalisado++) {

        // Se o valor analisado não é primo, eu pulo seu crivo.
        if (isPrime[valorAnalisado] == false) {
            continue;
        }

        // Se é primo, eu removo todos os próximos itens que são multiplos dele.
        for (int j = valorAnalisado+1; j < isPrime.size(); j++) {
            if (j % valorAnalisado == 0) {
                isPrime[j] = false;
            }
        }
    }

    auto tempoFinal = chrono::high_resolution_clock::now();

    int qtdPrimos = 0;

    // Imprimindo a lista de números primos e contando seu total.
    for (int i = 0; i < isPrime.size(); i++) {
        if (isPrime[i] == true) {
            cout << i << endl;
            qtdPrimos++;
        }
    }

    cout << "# de primos encontrados: " << qtdPrimos << endl;
    cout << "tempo: " << (std::chrono::duration_cast<std::chrono::nanoseconds>(tempoFinal - tempoInicial).count()) / 1000000.0 << " ms" << endl;

    return 0;
}