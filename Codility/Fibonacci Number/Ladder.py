'''
Source : https://app.codility.com/programmers/lessons/13-fibonacci_numbers/ladder/
Result : https://app.codility.com/demo/results/trainingCWN7F6-HJN/
'''

def fibo(n):
    res = [0]*(n+2)
    res[1] = 1
    for i in range(2, n+2):
        res[i] = res[i-1] + res[i-2]
    return res

def solution(A, B):
    result = []
    fib = fibo(max(A))
    print(fib)
    for i in range(len(A)):
        '''
        if using this, just get 75% score on performance
        result.append(fib[A[i]+1]%2**B[i])
        ''' 
        result.append(fib[A[i]+1] & ((1<<B[i])-1))
    return result
