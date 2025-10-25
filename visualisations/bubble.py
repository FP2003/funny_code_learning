import matplotlib.pyplot as plt
import numpy as np


amount = 100
lst = np.random.randint(0, 100, amount)
x = np.arange(0, amount, 1)

n = len(lst)
plt.figure("Bubble Sort")
for i in range(n):
    for j in range(0, n - i - 1):
        plt.bar(x, lst)
        plt.pause(0.005)
        plt.clf()
        if lst[j] > lst[j + 1]:
            lst[j], lst[j + 1] = lst[j + 1], lst[j]
            colors = ['C0'] * n
            colors[j] = 'C3'
            colors[j + 1] = 'C3'
            plt.bar(x, lst, color=colors)
            plt.pause(0.02)
            plt.clf()

plt.show()