# 257. Binary Tree Paths

Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:
```
   1
 /   \
2     3
 \
  5
```

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3


[257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil { return []string{} }
    // shared by dfs
    path := []int{}
    res := []string{}
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        if root.Left == nil && root.Right == nil {
            path = append(path, root.Val)
            res = append(res, sliceToString(path))
            // recover
            path = path[:len(path)-1]
            return
        }
        path = append(path, root.Val)
        dfs(root.Left)
        dfs(root.Right)
        path = path[:len(path)-1]
    }
    dfs(root)
    return res
}

func sliceToString(slice []int) string {
    if len(slice) == 0 { return "" }
    res := strconv.Itoa(slice[0])
    for i:=1; i<len(slice); i++ {
        res = res+"->"+strconv.Itoa(slice[i])
    }
    return res
}
```
