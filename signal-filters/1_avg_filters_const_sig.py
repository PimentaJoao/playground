import matplotlib.pyplot as plt
from random import uniform

def create_array(start, step, end):
    arr = []
    i = 0
    while (i*step)+start <= end:
        arr.append(round(i*step + start, 3))
        i = i+1
    return arr
    
x = create_array(0, 0.2, 100)
y_noisy_sig = []
y_bulk_avg = []
y_iter_avg = []

base_sig = 2

# creating noisy signal
noise_lvl = 0.3
for _ in range(len(x)):
    noise = uniform(-1, 1) * noise_lvl
    y_noisy_sig.append(base_sig + noise)

# average (usual method)
sum = 0
for val in y_noisy_sig:
    sum = sum + val
y_bulk_avg = [sum/len(y_noisy_sig)] * len(y_noisy_sig)

# average (iterative)
i = 0
for val in y_noisy_sig:
    # initial value
    if len(y_iter_avg) == 0:
        y_iter_avg.append(val)
        i = i + 1
        continue
    
    # subsequent average numbers
    old_avg = y_iter_avg[i-1]
    new_avg = ((i-1)/i) * old_avg + (1/i) * val
    y_iter_avg.append(new_avg)
    i = i + 1

plt.plot(x, y_noisy_sig, linewidth=0.7, label='sinal ruidoso')
plt.plot(x, y_bulk_avg, linewidth=2, color='r', label='média comum')
plt.plot(x, y_iter_avg, linewidth=2, color='g', label='média iterativa')

plt.annotate('%0.4f' % y_bulk_avg[-1], xy=(1, y_bulk_avg[-1]), xytext=(6, 5), 
             xycoords=('axes fraction', 'data'),  color='r', fontsize=12, textcoords='offset points')
plt.annotate('%0.4f' % y_iter_avg[-1], xy=(1, y_iter_avg[-1]), xytext=(6, -10), 
             xycoords=('axes fraction', 'data'),  color='g', fontsize=12, textcoords='offset points')

plt.scatter(x[-1], y_bulk_avg[-1], color='r')
plt.scatter(x[-1], y_iter_avg[-1], color='g')

plt.legend(loc="upper left")
plt.title("Atuação de filtros de média comum e iterativo sobre sinal ruidoso")
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
