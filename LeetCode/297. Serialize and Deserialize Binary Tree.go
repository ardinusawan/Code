// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "strings"
)

var (
    resList []string
    counter int
)

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{} 
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    resList = make([]string, 0)
    
    serializerDFS(root)
    return strings.Join(resList, ",")
}

func serializerDFS(root *TreeNode) {
    if root == nil {
        resList = append(resList, "N")
        return
    }
     
    resList = append(resList, strconv.FormatInt(int64(root.Val), 10))
    serializerDFS(root.Left)
    serializerDFS(root.Right)
    return
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    s := strings.Split(data, ",") 
    counter = 0
    result := deserDFS(s)
    return result
}

func deserDFS(vals []string) *TreeNode {
    if vals[counter] == "N" {
        counter+=1
        return nil
    }
    var result TreeNode
    if i, err := strconv.ParseInt(vals[counter], 10, 64); err == nil {
        result.Val = int(i)
    }
    counter+=1
    result.Left = deserDFS(vals)
    result.Right = deserDFS(vals)
    return &result
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
