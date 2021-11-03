// https://leetcode.com/problems/reorder-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode)  {
    pointer := head
    arr := []int{}
    // make an array of value of linked-list
    for pointer != nil {
        arr = append(arr, pointer.Val)
        pointer = pointer.Next
    }
    var newArray []int
    // create new array contains merged data of array with reversed array
    for i:=0; i<len(arr);i++ {
        newArray = append(newArray, arr[i])
        newArray = append(newArray, arr[len(arr)-1-i])
    }
    
    var newHead []ListNode
    // create array of ListNode with empty next
    for _, v := range newArray {
        newHead = append(newHead, ListNode{Val: v})
    }
    
    // assign next to next item in the array
    for i:=0;i<(len(newHead)-1)/2;i++{
        newHead[i].Next = &newHead[i+1]
    }
    
    // assign head to new head
    head.Val = newHead[0].Val
    head.Next = newHead[0].Next
}

