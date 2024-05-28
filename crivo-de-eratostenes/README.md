# Implementações (minhas) do Crivo de Eratóstenes

O algoritmo do Crivo de Eratóstenes (Sieve of Eratosthenes) consiste no uso de um vetor de tamanho N, contendo valores do tipo `bool` (um bitset) onde cada número natural é dado pela sua respectiva posição no vetor, tendo seu valor `true` se for um número primo ou `false` caso contrário.

## Criação do Vetor

O vetor é inicializado com o valor `true` em todas suas posições, ou seja, assume-se, de início, que todos os números são primos. Em seguida é feita a análise (crivo) de número primo por número primo, começando em 2, marcando todos os subsequentes números que sejam divisíveis pelo primo analisado, esses números são então marcados com `false` (não primo).

## Uso do Vetor

Os números de interesse (suficientemente pequenos, menores que N) podem ser analisados com ordem O(1) (acesso do vetor na posição do número de interesse).

## Visualizando o Funcionamento

Visualização do funcionamento do algoritmo do Crivo de Eratóstenes:

![Crivo de Eratostenes em GIF](https://upload.wikimedia.org/wikipedia/commons/8/8c/New_Animation_Sieve_of_Eratosthenes.gif)
