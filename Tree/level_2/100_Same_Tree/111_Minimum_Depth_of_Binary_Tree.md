# 111. Minimum Depth of Binary Tree

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    // make sure input is not nil
    if root == nil { return 0 }
    min := math.MaxInt32
    var dfs func(root *TreeNode, depth int)
    dfs = func(root *TreeNode, depth int) {
        if root == nil { return }
        if root.Left == nil && root.Right == nil {
            if min > depth { min = depth }
            return
        }
        
        // if depth >= min, we don't need to go further => cut dfs
        if depth >= min { return }
        
        dfs(root.Left, depth+1)
        dfs(root.Right, depth+1)
    }
    dfs(root, 1)
    return min
}
```
