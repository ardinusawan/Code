import sys
def fib(n):
    if n < 2:
        return n

    cache = [-1]*(n+1)
    cache[0] = 0
    cache[1] = 1

    return _fib(n, cache)


def _fib(n, cache):
    if cache[n] >= 0:
        return cache[n]

    cache[n] = _fib(n-1, cache) + _fib(n-2, cache)
    
    return cache[n]

if __name__ == "__main__":
    n = int(sys.argv[1])
    print(fib(n))

