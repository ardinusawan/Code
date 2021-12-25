// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    // preorder = [3,9,20,15,7]
    // inorder = [9,3,15,20,7]
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    val := preorder[0]
    var result TreeNode
    result.Val = val
    for i, v := range inorder {
        if v == val {
            // [1:7] -> inclusion -> get value from index 1 until 6
            // [:7] get value from beginning until index 6
            
            // preorder = [9]
            // inorder = [9]
            result.Left = buildTree(preorder[1:i+1], inorder[:i])
            
            // preorder = [20,15,7]
            // inorder = [15,20,7]
            result.Right = buildTree(preorder[i+1:], inorder[i+1:])
            break
        }
    }
    return &result
}
