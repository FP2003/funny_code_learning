import matplotlib.pyplot as plt
import numpy as np

amount = 100
lst = np.random.randint(0, 100, amount)
x = np.arange(amount)

n = len(lst)
plt.figure("Selection Sort")
for step in range(n):
    min_idx = step
    for i in range(step + 1, n):
        colors = ['C0'] * n
        colors[min_idx] = 'C1'   
        colors[i] = 'C2'
        plt.bar(x, lst, color=colors)
        plt.pause(0.005)
        plt.clf()

        if lst[i] < lst[min_idx]:
            min_idx = i

    if min_idx != step:
        lst[step], lst[min_idx] = lst[min_idx], lst[step]
        colors = ['C0'] * n
        colors[step] = 'C3'
        colors[min_idx] = 'C3'
        plt.bar(x, lst, color=colors)
        plt.pause(0.02)
        plt.clf()

plt.show()