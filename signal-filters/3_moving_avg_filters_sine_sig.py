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
y_bulk_mov_avg = []
y_iter_mov_avg = []

w = 10 # Tamanho da janela

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
        # reusa a média anterior e ajusta com o novo valor na janela
        y_iter_mov_avg.append(prev_avg + (measurement - y_noisy_sig[i - w]) / w)

plt.plot(x, y, linewidth=1.5, color='k', label='sinal real')
plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
# plt.plot(x, y_bulk_mov_avg, linewidth=2, color='r', label='média móvel comum')    # Comentado por não adicionar valor
                                                                                    # visual à demonstração do filtro,
                                                                                    # em comparação com sua versão
                                                                                    # iterativa.
plt.plot(x, y_iter_mov_avg, linewidth=1.7, color='g', label='média móvel iterativa')
plt.legend(loc="upper left")
plt.title(f"Atuação do filtro de média móvel iterativo sobre sinal dinâmico ruidoso - Janela de tamanho {w}")
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
