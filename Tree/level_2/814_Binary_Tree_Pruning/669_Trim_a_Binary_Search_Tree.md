# 669. Trim a Binary Search Tree

Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in \[L, R\] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:
```
Input: 
    1
   / \
  0   2

  L = 1
  R = 2

Output: 
    1
      \
       2
Example 2:
Input: 
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output: 
      3
     / 
   2   
  /
 1
```
[669. Trim a Binary Search Tree](https://leetcode.com/problems/trim-a-binary-search-tree/)


Recursion: consider that the recursive function itself is correct, then consider internal logic and boundary conditions.

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil 
	}

	// root.Val < L => we left the right part
	if root.Val < L {
        return trimBST(root.Right, L, R)
	}

	// root.Val > L => we left the left part
	if root.Val > R {
        return trimBST(root.Left, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
```
