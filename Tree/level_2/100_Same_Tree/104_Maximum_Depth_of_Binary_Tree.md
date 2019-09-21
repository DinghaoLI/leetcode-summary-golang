# 104. Maximum Depth of Binary Tree

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

[104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    max := 0 
    var dfs func(root *TreeNode, depth int) 
    dfs = func(root *TreeNode, depth int) {
        if root == nil { return }
        if depth+1 > max { max = depth+1 }
        dfs(root.Left, depth+1)
        dfs(root.Right, depth+1)
    } 
    dfs(root, 0)
    return max
}
```
