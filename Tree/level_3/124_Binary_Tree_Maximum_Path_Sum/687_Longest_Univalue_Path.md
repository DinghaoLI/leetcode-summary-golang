# 687. Longest Univalue Path

Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.

The length of path between two nodes is represented by the number of edges between them.

 

Example 1:

Input:
```
              5
             / \
            4   5
           / \   \
          1   1   5
```
Output: 2

 

Example 2:

Input:
```
              1
             / \
            4   5
           / \   \
          4   4   5
```
Output: 2

 

Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.


[687. Longest Univalue Path](https://leetcode.com/problems/longest-univalue-path/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
    maxLen := 0
    // dfs will replace maxLen (if root can connect left leaf and right leaf) by higher value
    // and return the max univalue path which start from root
    var dfs func(root *TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil { return 0 }
        
        left := dfs(root.Left)
        right := dfs(root.Right)
        res := 0
        
        // path for left leaf
        if root.Left != nil && root.Val == root.Left.Val {
            maxLen = max(maxLen, left+1)
            res = max(res, left+1)
        }
        
        // path for right leaf
        if root.Right != nil && root.Val == root.Right.Val {
            maxLen = max(maxLen, right+1)
            res = max(res, right+1)
        }
        
        // if root can connect both right side and left side
        if root.Right != nil && root.Val == root.Right.Val &&
           root.Left != nil && root.Val == root.Left.Val {
            maxLen = max(maxLen, right+left+2)
        }
        return res
    }
    dfs(root)
    return maxLen
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
