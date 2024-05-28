public class Crivo {

    public static void main(String[] args) {
        
        // Espaço analisado.
        int valorLimite = 10000000;

        // Raíz quadrada do valor limite analisado, determinando o maior número a ser analisado.
        int multiploLimite = (int) Math.floor(Math.sqrt(valorLimite));

        boolean[] isPrime = new boolean[valorLimite];

        // Considero que todos os números são primos, inicialmente.
        for (int i = 0; i < isPrime.length; i++) {
            isPrime[i] = true;
        }

        long tempoInicial = System.nanoTime();

        // Otimização: Removo números que eu sei que não são primos:
        // 0, 1 e todos os pares que não sejam o 2.
        isPrime[0] = false;
        isPrime[1] = false;
        for (int i = 4; i < isPrime.length; i+=2) {
            isPrime[i] = false;
        }

        // Começo a rodar o algoritmo do crivo à partir do primeiro primo ímpar conhecido (3).
        for (int valorAnalisado = 3; valorAnalisado < multiploLimite; valorAnalisado+=2) {

            // Se o valor analisado não é primo, eu pulo seu crivo.
            if (isPrime[valorAnalisado] == false) {
                continue;
            }

            // Se é primo, eu removo todos os próximos itens ímpares que são multiplos dele.
            for (int j = valorAnalisado+2; j < isPrime.length; j+=2) {
                if (j % valorAnalisado == 0) {
                    isPrime[j] = false;
                }
            }
        }

        long tempoFinal = System.nanoTime();

        int qtdPrimos = 0;

        // Imprimindo a lista de números primos e contando seu total.
        for (int i = 0; i < isPrime.length; i++) {
            if (isPrime[i] == true) {
                System.out.println(i);
                qtdPrimos++;
            }
        }

        System.out.println("# de primos encontrados: " + qtdPrimos);
        System.out.println("tempo: " + (tempoFinal - tempoInicial) / 1_000_000.0 + " ms");
    }
}