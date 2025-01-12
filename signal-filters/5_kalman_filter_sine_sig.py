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
y_kalman = []

amplitude = 0.2  # Amplitude da onda senoidal
frequency = 0.03  # Frequência da onda senoidal
base_sig = 2

# Criando sinal ruidoso
noise_lvl = 0.05
for t in x:
    noise = uniform(-1, 1) * noise_lvl
    sine_wave = base_sig + amplitude * math.sin(2 * math.pi * frequency * t)
    y.append(sine_wave)
    y_noisy_sig.append(sine_wave + noise)

# Parâmetros do Filtro de Kalman
A = 1  # Modelo do sistema
H = 1  # Modelo de medição
Q = 1e-5  # Variância do processo (ajustado para filtrar ruídos pequenos)
R = 0.001  # Variância do ruído de medição

# Estados iniciais
x_est = base_sig  # Estimativa inicial (próxima ao valor médio do sinal base)
P = 1  # Covariância inicial

# Implementação do Filtro de Kalman
for measurement in y_noisy_sig:
    # Previsão (Prediction)
    x_pred = A * x_est
    P_pred = A * P * A + Q

    # Ganho de Kalman
    K = P_pred * H / (H * P_pred * H + R)

    # Atualização (Update)
    x_est = x_pred + K * (measurement - H * x_pred)
    P = P_pred - K * H * P_pred

    # Salvar resultado
    y_kalman.append(x_est)

plt.plot(x, y, linewidth=1.5, color='k', label='sinal real')
plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
plt.plot(x, y_kalman, linewidth=1.7, color='g', label='kalman filter')
plt.legend(loc="upper left")
plt.title(f"Atuação do filtro de Kalman sobre sinal dinâmico ruidoso")
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
