/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
    result int
)

func maxPathSum(root *TreeNode) int {
    result = root.Val
    dfs(root)
    return result
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    // if value < 0 dont add it up
    lMax := int(math.Max(float64(dfs(root.Left)), float64(0)))
    rMax := int(math.Max(float64(dfs(root.Right)), float64(0)))
    
    // root + left + right path
    result = int(math.Max(float64(result), float64(root.Val + lMax + rMax)))
    
    // max of root + left || root + right
    // because it dfs, only from one path
    return root.Val + int(math.Max(float64(lMax), float64(rMax)))
}
