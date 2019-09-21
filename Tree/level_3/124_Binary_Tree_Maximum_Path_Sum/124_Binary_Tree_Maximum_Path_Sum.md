# 124. Binary Tree Maximum Path Sum

Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]
```
       1
      / \
     2   3
```
Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]
```
   -10
   / \
  9  20
    /  \
   15   7
```
Output: 42


[124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    if root == nil { return 0 }
    maxValue := root.Val
    
    // dfs return the max sum path that starts from root
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil { return 0 }
        left := max(0, dfs(root.Left))
        right := max(0, dfs(root.Right))
        sum := root.Val + left + right
        if sum > maxValue {
            maxValue = sum
        }
        
        return max(left, right) + root.Val
    }

    dfs(root)
    return maxValue
}

func max(a,b int) int {
    if a > b { return a }
    return b
}
```
