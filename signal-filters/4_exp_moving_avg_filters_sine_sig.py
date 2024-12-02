import matplotlib.pyplot as plt
import math
from random import uniform

def create_array(start, step, end):
    arr = []
    i = 0
    while (i*step)+start <= end:
        arr.append(round(i*step + start, 3))
        i = i+1
    return arr
    
x = create_array(0, 0.2, 100)
y = []
y_noisy_sig = []
y_exp_mov_avg = []

alpha = 0.2

amplitude = 0.2  # Amplitude da onda senoidal
frequency = 0.05  # Frequência da onda senoidal
base_sig = 2

# Criando sinal ruidoso
noise_lvl = 0.2
for t in x:
    noise = uniform(-1, 1) * noise_lvl
    sine_wave = base_sig + amplitude * math.sin(2 * math.pi * frequency * t)
    y.append(sine_wave)
    y_noisy_sig.append(sine_wave + noise)

# Média móvel exponencial
for i, measurement in enumerate(y_noisy_sig):
    if i == 0:
        y_exp_mov_avg.append(measurement)
    else:
        y_exp_mov_avg.append((1 - alpha) * y_exp_mov_avg[-1] + alpha * measurement)

plt.plot(x, y, linewidth=1.5, color='k', label='sinal real')
plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
plt.plot(x, y_exp_mov_avg, linewidth=1.7, color='g', label='média móvel exponencial')
plt.legend(loc="upper left")
plt.title(f"Atuação do filtro de média móvel exponencial sobre sinal dinâmico ruidoso - Alfa de {alpha}")
plt.ylabel("Altura (m)")
plt.xlabel("tempo (s)")
plt.ylim((1.5, 2.5))
plt.subplots_adjust(
    top=0.93,
    bottom=0.11,
    left=0.075,
    right=0.94,
    hspace=0.2,
    wspace=0.2
)
plt.show()
