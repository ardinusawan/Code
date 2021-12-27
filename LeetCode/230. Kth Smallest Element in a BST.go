// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
    counter int
    smallest int
)

type visit map[*TreeNode]bool

func kthSmallest(root *TreeNode, k int) int {
    // recursive inorder = dfs from left
    // inorder = left child, parent, right
    // doing inorder make we visiting node with sorted value
    
    counter = 0
    smallest = 10001
    visited := make(visit)
    dfs(root, k, visited)
    return smallest
}

func dfs(root *TreeNode, target int, visited visit) int {
    if root == nil {
        return smallest
    }
    
    // visit unvisited node only and if target is not reached
    
    if root.Left != nil && visited[root.Left] == false && len(visited) < target {
        dfs(root.Left, target, visited)
    }
    
    visited[root] = true
    if len(visited) <= target {
        smallest = root.Val
    }
        
    if len(visited) == target {
        return smallest
    }
    
    if root.Right != nil && visited[root.Right] == false && len(visited) < target {
        dfs(root.Right, target, visited)
    }
    
    return smallest
}
