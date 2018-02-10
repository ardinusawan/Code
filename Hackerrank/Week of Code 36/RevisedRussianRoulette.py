# Source : https://www.hackerrank.com/contests/w36/challenges/revised-russian-roulette

#!/bin/python3

import sys

def revisedRussianRoulette(doors):
    minimal=0
    maximal=0
    for index in range(len(doors)):
        if(doors[index]==1):
            doors[index] =0
            minimal+=1
            if(index!=len(doors)-1 and doors[index+1]==1):
                doors[index+1]=0
                maximal+=2
            
            else:
                maximal+=1
        else:
            continue
        #elif(doors[index]==1 and doors[index+1]==0):
        #    minimal+=1
        #    maximal+=1
    return minimal, maximal
        
    # Complete this function

if __name__ == "__main__":
    n = int(input().strip())
    doors = list(map(int, input().strip().split(' ')))
    result = revisedRussianRoulette(doors)
    print (" ".join(map(str, result)))


