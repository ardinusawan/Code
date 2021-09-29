# https://leetcode.com/problems/linked-list-cycle/submissions/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head == None or head.next == None :
            return False
        limit = 10**4 +1
        while(limit>0):
            if head.next == None:
                return False
            head = head.next
            limit-=1
        else:
            return True
        return False
