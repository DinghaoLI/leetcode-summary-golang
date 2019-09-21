# 129. Sum Root to Leaf Numbers

Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

Note: A leaf is a node with no children.

Example:

Input: [1,2,3]
```
    1
   / \
  2   3
```
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
Example 2:

Input: [4,9,0,5,1]
```
    4
   / \
  9   0
 / \
5   1
```
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.


[129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    if root == nil { return 0 }
    // shared by dfs
    path := []int{}
    res := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        if root.Left == nil && root.Right == nil {
            path = append(path, root.Val)
            res += sliceToInt(path)
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

func sliceToInt(slice []int) int {
    if len(slice) == 0 { return 0 }
    res := slice[0]
    for i:=1; i<len(slice); i++ {
        res = res*10 + slice[i] 
    }
    return res
}
```
