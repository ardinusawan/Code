# https://www.hackerrank.com/challenges/stockmax/problem

#!/bin/python3

import sys

def stockmax(prices):
    # Complete this function
    profit = 0
    maxx = prices[-1]
    
    for i, val in reversed(list(enumerate(prices))):
        if maxx > val:
            profit+=maxx-val
        elif maxx < val:
            maxx = val
    return profit

if __name__ == "__main__":
    t = int(input().strip())
    for a0 in range(t):
        n = int(input().strip())
        prices = list(map(int, input().strip().split(' ')))
        result = stockmax(prices)
        print(result)


