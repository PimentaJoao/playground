# Maior número

Algoritmo para encontrar o maior número de um grande vetor de números naturais, dividindo-o igualmente para cada uma das 12 *threads* disponíveis no meu computador atualmente.

## Resultados

Algoritmo puramente sequencial:

~3245 ms

Algoritmo paralelizado com openmp (+87,82%):

~395 ms

## Execução

Algoritmo puramente sequencial:

```bash
g++ maior_numero.cpp && ./a.out
```

Algoritmo paralelizado com OpenMP:

```bash
g++ -fopenmp maior_numero_openmp.cpp && ./a.out
```

### Possíveis melhorias

- É interessante explorar outras otimizações de compilação do g++;
- Explorar as outras diretivas de compilação oferecidas pelo OpenMP.
