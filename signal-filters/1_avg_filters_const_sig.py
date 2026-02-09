from random import uniform
import utils

dt = 0.2 # Taxa de amostragem
n = 100 # Número de amostras

# Eixo x do gráfico
x = utils.create_array(0, dt, n)

# Eixo y do gráfico
y_noisy_sig = []
y_bulk_avg = []
y_iter_avg = []

base_sig = 2 # Metros
noise_lvl = 0.3 # Ruído de +- 30 centímetros

for _ in range(len(x)):
    noise = uniform(-1, 1) * noise_lvl
    y_noisy_sig.append(base_sig + noise)

# Filtro de média (método comum)
sum = 0
for val in y_noisy_sig:
    sum = sum + val
y_bulk_avg = [sum/len(y_noisy_sig)] * len(y_noisy_sig)

# Filtro de média (iterativa)
for i, y_val in enumerate(y_noisy_sig):
    # valor inicial
    if len(y_iter_avg) == 0:
        y_iter_avg.append(y_val)
        continue
    
    # valores subsequentes
    old_avg = y_iter_avg[i-1]
    new_avg = ((i-1)/i) * old_avg + (1/i) * y_val
    y_iter_avg.append(new_avg)

utils.plot_avg_filters_const_sig(x, y_noisy_sig, y_bulk_avg, y_iter_avg)
