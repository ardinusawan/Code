from math import sqrt

def solution(n, p, q):
    # Find out all the primes with Sieve of Eratosthenes
    prime_table = [False]*2+[True]*(n-1)
    prime=[]
    i=2
    while(i*i<=n):
        k=i*i
        while(k<=n):
            if(prime_table[k]):
                prime_table[k]=False
            k+=i
        i+=1
    for element in range(0,n+1):
        if(prime_table[element]):
            prime.append(element)
    prime_count=len(prime)

    # Compute the semiprimes information
    semiprime=[0] * (n+1)
    # Find out all the semiprimes.
    # semiprime[i] == 1 when i is semiprime, or
    #                 0 when i is not semiprime.
    for x in range(0, prime_count):
        for y in range(x, prime_count):
            product=prime[x]*prime[y]
            if product>n:
                break
            semiprime[product]=1
    
    # Compute the number of semiprimes until each position.
    # semiprime[i] == k means:
    # in the range (0,i] there are k semiprimes.
    for index in range(1, n+1):
        semiprime[index] += semiprime[index-1]
        
    # the number of semiprimes within the range [ P[K], Q[K] ]
    # should be semiprime[Q[K]] - semiprime[P[K]-1]
    question_len = len(p)
    result = [0]*question_len
    for index in range(question_len):
        result[index] = semiprime[q[index]] - semiprime[p[index]-1]
    return result
    
