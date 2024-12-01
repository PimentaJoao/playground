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
y_bulk_avg = []
y_iter_avg = []

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

# Média (método comum)
sum = 0
for val in y_noisy_sig:
    sum = sum + val
y_bulk_avg = [sum/len(y_noisy_sig)] * len(y_noisy_sig)

# Média (iterativa)
i = 0
for val in y_noisy_sig:
    # valor inicial
    if len(y_iter_avg) == 0:
        y_iter_avg.append(val)
        i = i + 1
        continue
    
    # valores subsequentes
    old_avg = y_iter_avg[i-1]
    new_avg = ((i-1)/i) * old_avg + (1/i) * val
    y_iter_avg.append(new_avg)
    i = i + 1

plt.plot(x, y, linewidth=2, color='k', label='sinal real')
plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
plt.plot(x, y_bulk_avg, linewidth=2, color='r', label='média comum')
plt.plot(x, y_iter_avg, linewidth=2, color='g', label='média iterativa')
plt.legend(loc="upper left")
plt.title("Atuação de filtros de média comum e iterativo sobre sinal dinâmico ruidoso")
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
