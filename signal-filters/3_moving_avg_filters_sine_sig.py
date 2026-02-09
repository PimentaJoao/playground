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
y_bulk_mov_avg = []
y_iter_mov_avg = []

w = 15 # Tamanho da janela
amplitude = 0.2 # Amplitude da onda senoidal (20 centímetros)
frequency = 0.05 # Frequência da onda senoidal
base_sig = 2 # Metros
noise_lvl = 0.2 # Ruído de +- 20 centímetros

for t in x:
    noise = uniform(-1, 1) * noise_lvl
    sine_wave = base_sig + amplitude * math.sin(2 * math.pi * frequency * t)
    y_groundtruth.append(sine_wave)
    y_noisy_sig.append(sine_wave + noise)

# Média móvel (método comum)
for i, measurement in enumerate(y_noisy_sig):
    if i < w:
        # enquanto não houver w elementos para o filtro, calcular a média dos disponíveis
        y_bulk_mov_avg.append(sum(y_noisy_sig[:i + 1]) / (i + 1))
    else:
        y_bulk_mov_avg.append(sum(y_noisy_sig[i - w + 1:i + 1]) / w)

# Média móvel (iterativo)
for i, measurement in enumerate(y_noisy_sig):
    if i == 0:
        y_iter_mov_avg.append(measurement)
        continue

    # atualiza a média de forma incremental
    prev_avg = y_iter_mov_avg[-1]
    if i < w:
        # enquanto não houver w elementos, considerar todos disponíveis
        y_iter_mov_avg.append(prev_avg + (measurement - prev_avg) / (i + 1))
    else:
        # reutiliza a média anterior e ajusta com o novo valor na janela
        y_iter_mov_avg.append(prev_avg + (measurement - y_noisy_sig[i - w]) / w)

utils.plot_moving_avg_filters_sine_sig(x, y_groundtruth, y_noisy_sig, y_bulk_mov_avg, y_iter_mov_avg, w)
