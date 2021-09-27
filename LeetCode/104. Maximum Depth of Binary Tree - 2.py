# https://leetcode.com/problems/maximum-depth-of-binary-tree/submissions/

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    maxD = 0
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        self.dfs(root, 0)
        return self.maxD
    
    def dfs(self, node, leaf):
        if node == None:
            return
        newL = leaf+1
        self.maxD = max(self.maxD, newL)
        if node.left!= None:
            self.dfs(node.left, newL)
        if node.right!= None:
            self.dfs(node.right, newL)
        return
