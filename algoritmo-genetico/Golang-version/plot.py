import pandas as pd
import matplotlib.pyplot as plt

# Carrega o arquivo CSV utilizando o pandas
df = pd.read_csv('~/Github/playground/algoritmo-genetico/Golang-version/dados.csv')

# Extrai os valores das colunas Best, Average e Worst para cada linha
iterations = df.index.tolist()  # Obtém a lista de indices como iteracões
best_values = df['Best'].tolist()
average_values = df['Average'].tolist()
worst_values = df['Worst'].tolist()

# Plota os valores em um gráfico de linhas
plt.figure(figsize=(10, 6))

# Plotando cada série de dados como uma linha
plt.plot(iterations, best_values, linestyle='-', color='green', label='Best')
plt.plot(iterations, average_values, linestyle='-', color='blue', label='Average')
plt.plot(iterations, worst_values, linestyle='-', color='red', label='Worst')

# Adicionando titulo e rotulos aos eixos
plt.title('Values Comparison Over Iterations')
plt.xlabel('Iteration')
plt.ylabel('Value')

# Adicionando legenda
plt.legend()

# Adicionando grade ao gráfico
plt.grid(True)

# Salva o gráfico como um arquivo PNG
plt.savefig('grafico_linhas.png', bbox_inches='tight')

# Mostra o gráfico plotado
plt.show()