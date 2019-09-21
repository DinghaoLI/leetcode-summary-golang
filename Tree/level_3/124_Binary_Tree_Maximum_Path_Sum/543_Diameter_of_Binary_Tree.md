# 543. Diameter of Binary Tree

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree
```
          1
         / \
        2   3
       / \     
      4   5   
``` 
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path b


[543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/s)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    d := 0
    
    // dfs calculate max diameter from root
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil { return 0 }
        left := dfs(root.Left)
        right := dfs(root.Right)
        sum := left + right
        if sum > d {
            d = sum
        }
        
        return max(left, right)+1
    }
    dfs(root)
    return d
}

func max(a,b int) int {
    if a > b { return a }
    return b
}
```
