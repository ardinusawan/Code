// https://leetcode.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var dummy ListNode
    // tail will like pointer. It will move from beginning until last node.
    // we use address of dummy to start the tail so dummy can follow the path without moving
    tail := &dummy
    
    for list1 != nil && list2 != nil {
        // next value is lowest head in two linked list
        // and then move head to next node
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        // move tail to next node
        tail = (*tail).Next
    }
    
    // if len is not same, get latest data to link list
    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }
    return dummy.Next
}
