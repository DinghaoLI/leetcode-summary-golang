# 113. Path Sum II

Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,
```
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
```
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]

[113. Path Sum II](https://leetcode.com/problems/path-sum-ii/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    var result [][]int
    var solution []int
    var preOrder func(root *TreeNode, sum int)
    preOrder = func(root *TreeNode, sum int) {
        if root == nil { return }
        rest := sum - root.Val        
        if rest == 0 && root.Left == nil && root.Right == nil {
            solution = append(solution, root.Val)
            tmp := make([]int, len(solution))
            copy(tmp, solution)
            result = append(result, tmp)
            solution = solution[:len(solution)-1]
            return
        }
        solution = append(solution, root.Val)
        preOrder(root.Left, rest)
        preOrder(root.Right, rest)
        solution = solution[:len(solution)-1]
    }
    preOrder(root, sum)
    return result
}
```

