# 965. Univalued Binary Tree

A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.

[965. Univalued Binary Tree](https://leetcode.com/problems/univalued-binary-tree/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil { return true }
    value := root.Val
    var dfs func(root *TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil { return true }
        return root.Val == value && dfs(root.Left) && dfs(root.Right)
    }
    return dfs(root)
}
```
