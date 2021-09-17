# https://leetcode.com/problems/maximum-depth-of-binary-tree/

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def maxDepth(self, root: 'TreeNode') -> 'int':
        if not root:
            return 0
        
        lMaxDepth = self.maxDepth(root.left)
        rMaxDepth = self.maxDepth(root.right)
        if lMaxDepth > rMaxDepth:
            return lMaxDepth + 1
        else:
            return rMaxDepth + 1
        
        
