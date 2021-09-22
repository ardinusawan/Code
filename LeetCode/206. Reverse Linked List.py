# https://leetcode.com/problems/reverse-linked-list/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head
        prevNode = ListNode(head.val)
        curr = head.next
        while(curr != None):
            prevNode = ListNode(curr.val, prevNode)
            curr = curr.next
        return prevNode
