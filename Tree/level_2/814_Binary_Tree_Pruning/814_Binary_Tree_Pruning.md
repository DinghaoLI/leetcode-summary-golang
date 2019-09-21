# 814. Binary Tree Pruning

We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1.

Return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

(Recall that the subtree of a node X is X, plus every node that is a descendant of X.)

Example 1:
Input: [1,null,0,0,1]
Output: [1,null,0,null,1]
 
Explanation: 
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.


[814. Binary Tree Pruning](https://leetcode.com/problems/binary-tree-pruning/)

Recursively consider the following situation:

root: nil or Val == 1 or Val == 0

leafs: pruneTree(root.Left) == nil && pruneTree(root.Right) == nil ?


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    
    if root.Val == 0 {
        if root.Left == nil && root.Right == nil {
            return nil
        }
    }
    return root
}
```
