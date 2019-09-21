# 101. Symmetric Tree

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }
    var symm func(p *TreeNode, q *TreeNode) bool
    symm = func(p *TreeNode, q *TreeNode) bool {
        if p == nil && q == nil { return true }
        if p == nil || q == nil { return false }
        
        return p.Val == q.Val && symm(p.Left, q.Right) && symm(p.Right, q.Left)
    }
    
    return symm(root.Left, root.Right)
}

```
