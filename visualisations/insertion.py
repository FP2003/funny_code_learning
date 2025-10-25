import matplotlib.pyplot as plt
import numpy as np

amount = 100
lst = np.random.randint(0, 100, amount)
x = np.arange(amount)

n = len(lst)
plt.figure("Insertion Sort")
for step in range(1, n):
    key = lst[step]
    j = step - 1
    
    colors = ['C0'] * n
    colors[step] = 'C1'
    plt.bar(x, lst, color=colors)
    plt.pause(0.005)
    plt.clf()
    
    while j >= 0 and key < lst[j]:
        lst[j + 1] = lst[j]
        j -= 1
        colors = ['C0'] * n
        colors[j + 1] = 'C2'
        colors[step] = 'C1'
        plt.bar(x, lst, color=colors)
        plt.pause(0.005)
        plt.clf()
        
    lst[j + 1] = key
    colors = ['C0'] * n
    colors[j + 1] = 'C3'
    plt.bar(x, lst, color=colors)
    plt.pause(0.02)
    plt.clf()

plt.show()
    
        
