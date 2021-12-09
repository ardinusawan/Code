// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // sum node on the list
    dummy := head
    sum := 0
    for dummy != nil {
        sum+=1
        dummy = dummy.Next
    }
    
    // edge case
    if sum == n {
        return head.Next
    }
    
    // p as address of head
    p := head
    i := 1
    for head != nil {
        // lets say sum = 5
        // i = 3
        // n = 2, so 5-3 == 2. Skip next node
        if sum-i == n {
            p.Next = p.Next.Next
            break
        } else {
            // keep walking
            p = (*p).Next
        }
        i+=1
    }
    
    return head
}
