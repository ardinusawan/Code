# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def gcd(a, b):
    if a % b == 0:
        return b
    else:
        return gcd(b, a % b)

def solution(N, M):
    # find lcm
    res_gcd = gcd(N, M)
    lcm = (N)/res_gcd
    return int(lcm)
