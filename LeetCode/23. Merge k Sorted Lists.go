// https://leetcode.com/problems/merge-k-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    // split n list to just two and merged those. Repeat
    // just like merge search
    for len(lists) > 1 {
        var mergedList []*ListNode
        for i:=0;i<len(lists);i+=2 {
            l1 := lists[i]
            var l2 *ListNode
            if i+1 < len(lists) {
                l2 = lists[i+1]
            }
            mergedList = append(mergedList, mergeList(l1, l2))
        }
        lists = mergedList
    }
    return lists[0]
}

func mergeList(l1, l2 *ListNode) *ListNode {
    var dummy ListNode
    // tail will like pointer. It will move from beginning until last node.
    // we use address of dummy to start the tail so dummy can follow the path without moving
    tail := &dummy
    
    for l1 != nil && l2 != nil {
        // next value is lowest head in two linked list
        // and then move head to next node
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        // move tail to next node
        tail = (*tail).Next
    }
    
    // if len is not same, get latest data to link list
    if l1 != nil {
        tail.Next = l1
    } else if l2 != nil {
        tail.Next = l2
    }
    return dummy.Next
}
