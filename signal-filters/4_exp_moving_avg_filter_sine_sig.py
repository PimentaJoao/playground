import math
import utils
from random import uniform

dt = 0.2 # Taxa de amostragem
n = 100 # Número de amostras

# Eixo x do gráfico
x = utils.create_array(0, dt, n)

# Eixos y dos gráficos
y_groundtruth = []
y_noisy_sig = []
y_exp_mov_avg = []

alpha = 0.1
amplitude = 0.2 # Amplitude da onda senoidal (20 centímetros)
frequency = 0.05 # Frequência da onda senoidal
base_sig = 2 # Metros
noise_lvl = 0.2 # Ruído de +- 20 centímetros

for t in x:
    noise = uniform(-1, 1) * noise_lvl
    sine_wave = base_sig + amplitude * math.sin(2 * math.pi * frequency * t)
    y_groundtruth.append(sine_wave)
    y_noisy_sig.append(sine_wave + noise)

# Média móvel exponencial
for i, measurement in enumerate(y_noisy_sig):
    if i == 0:
        y_exp_mov_avg.append(measurement)
    else:
        y_exp_mov_avg.append((1 - alpha) * y_exp_mov_avg[-1] + alpha * measurement)

utils.plot_exp_moving_avg_filter_sine_sig(x, y_groundtruth, y_noisy_sig, y_exp_mov_avg, alpha)
