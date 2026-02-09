import matplotlib.pyplot as plt

# create_array creates values for plotting graphs.
#
# Example:
# 
# >>> x = create_array(0, 0.2, 100)
# >>> x
# [0, 0.2, 0.4, 0.6, ..., 99.6, 99.8, 100.0]
# 
# (this was made before I learned about numpy's linspace)
# 
def create_array(start, step, end):
    arr = []
    i = 0
    while (i*step)+start <= end:
        arr.append(round(i*step + start, 3))
        i = i+1
    return arr

def plot_avg_filters_const_sig(x, y_noisy_sig, y_bulk_avg, y_iter_avg):
    plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
    plt.plot(x, y_bulk_avg, linewidth=2, color='r', label='média comum')
    plt.plot(x, y_iter_avg, linewidth=2, color='g', label='média iterativa')
    plt.annotate('%0.4f' % y_bulk_avg[-1], xy=(1, y_bulk_avg[-1]), xytext=(6, 5), xycoords=('axes fraction', 'data'),  color='r', fontsize=12, textcoords='offset points')
    plt.annotate('%0.4f' % y_iter_avg[-1], xy=(1, y_iter_avg[-1]), xytext=(6, -10), xycoords=('axes fraction', 'data'),  color='g', fontsize=12, textcoords='offset points')
    plt.scatter(x[-1], y_bulk_avg[-1], color='r')
    plt.scatter(x[-1], y_iter_avg[-1], color='g')
    plt.legend(loc="upper left")
    plt.title("Atuação de filtros de média comum e iterativo sobre sinal ruidoso")
    plt.ylabel("Altura (m)")
    plt.xlabel("tempo (s)")
    plt.ylim((1.5, 2.5))
    plt.subplots_adjust(top=0.93, bottom=0.11, left=0.075, right=0.94, hspace=0.2, wspace=0.2)
    plt.show()

def plot_avg_filters_sine_sig(x, y_groundtruth, y_noisy_sig, y_bulk_avg, y_iter_avg):
    plt.plot(x, y_groundtruth, linewidth=2, color='k', label='sinal real')
    plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
    plt.plot(x, y_bulk_avg, linewidth=2, color='r', label='média comum')
    plt.plot(x, y_iter_avg, linewidth=2, color='g', label='média iterativa')
    plt.legend(loc="upper left")
    plt.title("Atuação de filtros de média comum e iterativo sobre sinal dinâmico ruidoso")
    plt.ylabel("Altura (m)")
    plt.xlabel("tempo (s)")
    plt.ylim((1.5, 2.5))
    plt.subplots_adjust(top=0.93,bottom=0.11,left=0.075,right=0.94,hspace=0.2,wspace=0.2)
    plt.show()

def plot_moving_avg_filters_sine_sig(x, y_groundtruth, y_noisy_sig, y_bulk_mov_avg, y_iter_mov_avg, w):
    plt.plot(x, y_groundtruth, linewidth=1.5, color='k', label='sinal real')
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
    plt.subplots_adjust(top=0.93,bottom=0.11,left=0.075,right=0.94,hspace=0.2,wspace=0.2)
    plt.show()

def plot_exp_moving_avg_filter_sine_sig(x, y_groundtruth, y_noisy_sig, y_exp_mov_avg, alpha):
    plt.plot(x, y_groundtruth, linewidth=1.5, color='k', label='sinal real')
    plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
    plt.plot(x, y_exp_mov_avg, linewidth=1.7, color='g', label='média móvel exponencial')
    plt.legend(loc="upper left")
    plt.title(f"Atuação do filtro de média móvel exponencial sobre sinal dinâmico ruidoso - Alfa de {alpha}")
    plt.ylabel("Altura (m)")
    plt.xlabel("tempo (s)")
    plt.ylim((1.5, 2.5))
    plt.subplots_adjust(top=0.93,bottom=0.11,left=0.075,right=0.94,hspace=0.2,wspace=0.2)
    plt.show()

def plot_kalman_filter_1D_position_pos(t, y_groundtruth_pos, y_kalman_pos, y_noisy_pos):
    plt.plot(t, y_noisy_pos, linewidth=.8, label='sinal ruidoso')
    plt.plot(t, y_groundtruth_pos, linewidth=4, color='k', label='posição real')
    plt.plot(t, y_kalman_pos, linewidth=1.7, color='limegreen', label='estimativa do filtro de Kalman')
    plt.ylabel("Posição (m)")
    plt.xlabel("tempo (s)")
    plt.legend()
    plt.show()

def plot_kalman_filter_1D_position_vel(t, y_groundtruth_vel, y_kalman_vel):
    plt.plot(t, y_groundtruth_vel, linewidth=1.7, color='k', label='velocidade real')
    plt.plot(t, y_kalman_vel, linewidth=3, color='g', label='estimativa do filtro de Kalman')
    plt.ylabel("Velocidade (m/s)")
    plt.xlabel("tempo (s)")
    plt.legend()
    plt.show()

def plot_kalman_filter_1D_position_acc(t, y_groundtruth_acc, y_kalman_acc):
    plt.plot(t, y_groundtruth_acc, linewidth=1.7, color='k', label='aceleração real')
    plt.plot(t, y_kalman_acc, linewidth=3, color='g', label='estimativa do filtro de Kalman')
    plt.ylabel("Aceleração (m/s²)")
    plt.xlabel("tempo (s)")
    plt.legend()
    plt.show()

def plot_kalman_filter_1D_position_sigmas(t, sigmas):
    _, axes = plt.subplots(3, 3, figsize=(16, 12))
    axes = axes.flatten()

    for i in range(3):
        for j in range(3):
            idx = i * 3 + j
            label = f"Σ[{i},{j}]"
            values = [sigma[i, j] for sigma in sigmas]
            axes[idx].plot(t, values, linewidth=1.7, color='k', label=label)
            axes[idx].set_xlabel("tempo (s)")
            axes[idx].legend()
            axes[idx].grid(True)

    plt.tight_layout()
    plt.show()