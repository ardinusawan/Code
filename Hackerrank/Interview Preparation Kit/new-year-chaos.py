#https://www.hackerrank.com/challenges/new-year-chaos/problem
#https://ideone.com/cM0ZME

#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the minimumBribes function below.
def minimumBribes(q):
    q = [[i, v, 0] for i, v in enumerate(q)]

    result = 0
    swapped = False

    for k in range(len(q)-1):
        for i in range(len(q)-k-1):
            if q[i][1] > q[i+1][1]:
                q[i][2] += 1

                temp = q[i]
                q[i] = q[i+1]
                q[i+1] = temp

                result += 1
                swapped = True

            if q[i+1][2] > 2:
                print('Too chaotic')
                return

        if swapped:
            swapped = False
        else:
            break

    print(result)

if __name__ == '__main__':
    t = int(input())

    for t_itr in range(t):
        n = int(input())

        q = list(map(int, input().rstrip().split()))

        minimumBribes(q)

