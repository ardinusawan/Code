#!/bin/python3
# Source: https://www.hackerrank.com/challenges/minimum-swaps-2
# Inspired by: https://www.geeksforgeeks.org/minimum-number-swaps-required-sort-array/

import math
import os
import random
import re
import sys

# Complete the minimumSwaps function below.
def minimumSwaps(arr):
    n = len(arr)
    array_positions = [*enumerate(arr)]
    array_positions.sort(key = lambda v:v[1])
    visited = {k:False for k in range(n)}

    answer = 0
    for i in range(n): # i ==0, i == 1
        if visited[i] or array_positions[i][0] == i:
            continue

        cycle_size = 0
        j = i
        while not visited[j]:
            visited[j] = True

            j = array_positions[j][0]
            cycle_size +=1

        if cycle_size > 0:
            answer += (cycle_size - 1)
    return answer

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    res = minimumSwaps(arr)

    fptr.write(str(res) + '\n')

    fptr.close()
