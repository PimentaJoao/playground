import math
import utils
from random import uniform

dt = 0.2 # Taxa de amostragem
n = 100 # Número de amostras

# Eixo x do gráfico
x = utils.create_array(0, dt, n)

# Eixo y do gráfico
y_groundtruth = []
y_noisy_sig = []
y_bulk_avg = []
y_iter_avg = []

amplitude = 0.2 # Amplitude da onda senoidal (20 centímetros)
frequency = 0.05 # Frequência da onda senoidal
base_sig = 2 # Metros
noise_lvl = 0.2 # Ruído de +- 20 centímetros

for t in x:
    noise = uniform(-1, 1) * noise_lvl
    sine_wave = base_sig + amplitude * math.sin(2 * math.pi * frequency * t)
    y_groundtruth.append(sine_wave)
    y_noisy_sig.append(sine_wave + noise)

# Filtro de média (método comum)
sum = 0
for val in y_noisy_sig:
    sum = sum + val
y_bulk_avg = [sum/len(y_noisy_sig)] * len(y_noisy_sig)

# Filtro de média (iterativa)
for i, y_val in enumerate(y_noisy_sig):
    if i == 0:
        y_iter_avg.append(y_val)
    else:
        old_avg = y_iter_avg[i-1]
        new_avg = ((i-1)/i) * old_avg + (1/i) * y_val
        y_iter_avg.append(new_avg)

utils.plot_avg_filters_sine_sig(x, y_groundtruth, y_noisy_sig, y_bulk_avg, y_iter_avg)
