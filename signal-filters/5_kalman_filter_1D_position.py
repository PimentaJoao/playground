import numpy as np
import utils

dt = 0.2 # Taxa de amostragem
n = 150 # Número de amostras

groundtruth_initial_pos = 2  # Posição inicial real
groundtruth_initial_vel = -5 # Velocidade inicial real
groundtruth_acc = 0.8        # Aceleração (constante) real
noise_lvl = 10                # Ruído de +- 5 metros

# Eixo x do gráfico
t = np.linspace(0, (n - 1) * dt, n)

# Eixo y do gráfico
y_groundtruth_pos = groundtruth_initial_pos + groundtruth_initial_vel * t + 0.5 * groundtruth_acc * t**2
y_groundtruth_vel = groundtruth_initial_vel + groundtruth_acc * t
y_groundtruth_acc = np.ones(n) * groundtruth_acc
y_noisy_pos = y_groundtruth_pos + np.random.randn(n) * noise_lvl
y_kalman_pos = []
y_kalman_vel = []
y_kalman_acc = []

# Estimativa inicial
X = np.array([
    [0], # pos
    [0], # vel
    [0], # acc
])

# Matriz de transição de estados para o modelo matemático contendo a posição (pos), velocidade (vel) e aceleração (acc):
#
# pos: x = x0 + vΔt + (1/2)aΔt 
# vel: v = v0 + aΔt 
# acc: a = a0
# 
A = np.array([
    [1, dt, 0.5*dt**2],
    [0,  1, dt],
    [0,  0, 1]
])

# matriz de covariância "certeza geral do filtro"
Sigma = np.array([
    [1, 0, 0],
    [0, 1, 0],
    [0, 0, 1],
])

# matriz "máscara", que mapeia o formato (3x1) do estado para o formato medição (1x1)
C = np.array([
    [1, 0, 0]
])

# matriz de covariância "erro no modelo"
r = 0
R = np.array([
    [1, 0, 0],
    [0, 1, 0],
    [0, 0, 1],
]) * r

# matriz de covariância "erro da medição"
q = noise_lvl
Q = np.array([
    [1]
]) * q

Sigmas = []

# Kalman Filter loop
for i in range(n):
    Sigmas.append(Sigma)
    
    # Propagação
    X = A @ X
    Sigma = A @ Sigma @ A.T + R

    # Assimilação
    z = np.array([
        [y_noisy_pos[i]]
    ])
    K = Sigma @ C.T @ np.linalg.inv(C @ Sigma @ C.T + Q)
    X = X + K @ (z - C @ X)
    Sigma = (np.eye(3) - K @ C) @ Sigma


    y_kalman_pos.append(X[0][0])
    y_kalman_vel.append(X[1][0])
    y_kalman_acc.append(X[2][0])

# utils.plot_kalman_filter_1D_position_pos(t, y_groundtruth_pos, y_kalman_pos, y_noisy_pos)
# utils.plot_kalman_filter_1D_position_vel(t, y_groundtruth_vel, y_kalman_vel)
# utils.plot_kalman_filter_1D_position_acc(t, y_groundtruth_acc, y_kalman_acc)
utils.plot_kalman_filter_1D_position_sigmas(t, Sigmas)