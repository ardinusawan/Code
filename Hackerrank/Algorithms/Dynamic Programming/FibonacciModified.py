# https://www.hackerrank.com/challenges/fibonacci-modified/problem
#!/bin/python3

import sys
import math
def fibonacciModified(t1, t2, n):
    # Complete this function
    data = {}
    t = "t"
    base = 1
    
    data[t + str(base)] = t1
    data[t + str(base+1)] = t2

    for i in range(3, n+1):
        data[t+str(i)] = data[t+str(i-2)] + pow(data[t+str(i-1)],2)
    return (data[t+str(n)])
    

if __name__ == "__main__":
    t1, t2, n = input().strip().split(' ')
    t1, t2, n = [int(t1), int(t2), int(n)]
    result = fibonacciModified(t1, t2, n)
    print(result)


