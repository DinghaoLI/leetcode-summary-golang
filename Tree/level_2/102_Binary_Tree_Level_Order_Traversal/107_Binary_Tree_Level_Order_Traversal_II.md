# 107. Binary Tree Level Order Traversal II

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```
return its bottom-up level order traversal as:
```
[
  [15,7],
  [9,20],
  [3]
]
```
[107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    var res [][]int
    var preOrder func(root *TreeNode, depth int)
    preOrder = func(root *TreeNode, depth int) {
        if root == nil { return }
        if depth > len(res) {
            res = append([][]int{[]int{root.Val}}, res...)
        } else {
            res[len(res)-depth] = append(res[len(res)-depth], root.Val)
        }
        preOrder(root.Left, depth+1)
        preOrder(root.Right, depth+1)
    }
    preOrder(root, 1)
    return res
}
```
