# you can write to stdout for debugging purposes, e.g.
# print "this is a debug message"
from collections import deque

def solution(A, K):
    item = deque(A)
    item.rotate(K)
    return list(item)