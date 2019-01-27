#https://www.hackerrank.com/challenges/counting-valleys/

#!/bin/python3

import math
import os
import random
import re
import sys


# Complete the countingValleys function below.
def countingValleys(n, s):
    in_valey = 0
    sum_ = 0

    const = {
        'D': -1,
        'U': 1
    }
    for data in s:
        sum_ += const[data]

        if sum_ == -1 and const[data] == -1:
            in_valey += 1

    return in_valey



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    s = input()

    result = countingValleys(n, s)

    fptr.write(str(result) + '\n')

    fptr.close()

