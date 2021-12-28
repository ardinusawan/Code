// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

var found bool

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    found = false
    return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    
    // combination
    if (root == p || root.Left == p || root.Right == p) && (root == q || root.Left == q || root.Right == q) {
        found = true
    } else { // dont walk the tree if already found
        if root.Left != nil && found == false {
            return lowestCommonAncestor(root.Left, p, q)
        }
    
        if root.Right != nil && found == false {
            return lowestCommonAncestor(root.Right, p, q)
        }
    }
    return root
}
