# 110. Balanced Binary Tree

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

> a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

[110. Balanced Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var balanced func(root *TreeNode) int
    balanced = func(root *TreeNode) int {
        if root == nil { return 0 }
        left := balanced(root.Left)
        if left == -1 { return -1 }
        right := balanced(root.Right)
        if right == -1 { return -1 }
        if abs(left-right) > 1 { return -1 }
        return max(left, right)+1
    }
    return balanced(root) != -1
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}


func max(a, b int) int {
    if a < b { return b }
    return a
}
```
