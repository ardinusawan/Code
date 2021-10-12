# https://leetcode.com/problems/merge-two-sorted-lists/

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        h1 = l1
        h2 = l2
        
        if h1==None and h2==None:
            return None
        result = []
        result_n = ListNode() 
        while(h1 != None or h2 != None):
            if h1 != None:
                if h2 !=None and h1.val<=h2.val:
                    result.append(h1.val)
                    h1 = h1.next
                elif h2==None:
                    result.append(h1.val)
                    h1 = h1.next
            if h2 != None:
                if h1 !=None and h2.val<=h1.val:
                    result.append(h2.val)
                    h2 = h2.next
                elif h1==None:
                    result.append(h2.val)
                    h2 = h2.next
        rn = result_n
        for i in range(0, len(result)):
            result_n.val = result[i]
            if i < len(result)-1:
                result_n.next = ListNode(result[i+1])
            result_n = result_n.next
        return rn 
