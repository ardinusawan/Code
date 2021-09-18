# https://leetcode.com/problems/clone-graph/submissions/

"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        m = {}
        def dfs(node):
            if node.neighbors == []:
                return Node(node.val)
            if node in m:
                return m[node]
            copy = Node(node.val)
            m[node] = copy
            for neigh in node.neighbors:
                copy.neighbors.append(dfs(neigh))
            return copy 
        
        if node:
            return dfs(node)
        else:
            return node
        
