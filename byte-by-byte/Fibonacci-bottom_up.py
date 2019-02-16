import sys

def fib(n):
    if n == 0:
        return 0
    
    cache = [0]*(n+1)
    cache[1] = 1

    for i in range(2, len(cache)):
        cache[i] = cache[i-1] + cache[i-2]

    return cache[n] 

if __name__ == "__main__":
    n = int(sys.argv[1])
    print(fib(n))