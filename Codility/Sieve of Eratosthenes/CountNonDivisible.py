# Source : https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_non_divisible/

# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
from collections import Counter
from math import sqrt
def solution(A):
    '''
    Create sum of testcase for every element in testcase
    Create dict for divisors of every element
    Subtract every element in divisors dict with occurence with length of testcase array
    '''
    A_max = max(A)
    A_len = len(A)
    
    # count occurence data every testcase
    count = dict(Counter(A))
    
    # for every data have 1 as divisor
    divisors = {x:[1] for x in A}
    
    # find divisors less dan root of length A
    for divisor in range(2,int(sqrt(A_max))+1):
        multiple = divisor
        while multiple <= A_max:
            if multiple in divisors and not divisor in divisors[multiple]:
                divisors[multiple].append(divisor)
            multiple+=divisor
            
    # find divisors less dan root of length A
    for element in divisors:
        for data in divisors[element]:
            divisor = int(element/data)
            if divisor not in divisors[element]:
                divisors[element].append(divisor)
    
    result = []
    for element in A:
        result.append(A_len - sum(count.get(x,0) for x in divisors[element]))
    return result
