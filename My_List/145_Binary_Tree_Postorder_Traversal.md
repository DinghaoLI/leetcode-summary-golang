# 145. Binary Tree Postorder Traversal

Given a binary tree, return the postorder traversal of its nodes' values.

```
Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]

Follow up: Recursive solution is trivial, could you do it iteratively?
```

[145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// recursive
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    var post func(root *TreeNode) 
    post = func(root *TreeNode) {
        if root == nil { return }
        post(root.Left)
        post(root.Right)
        res = append(res, root.Val)
    }  
    post(root)
    return res
}

// iterative
func postorderTraversal(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
    cur := root
    var preVisited *TreeNode
	for len(stack) > 0 || cur != nil {
        // push left node to stack until the left leaf 
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        // pop a node
        node := stack[len(stack)-1]
        
        // if <node is leaf> || 
        // <node have not node.Right> ||
        // <we have visited node.Right> => we record this node
        if (node.Left == nil && node.Right == nil) || 
            node.Right == nil || preVisited == node.Right {
            // record
            result = append(result, node.Val)
            preVisited = node
            // remove it from stack
            stack = stack[:len(stack)-1]
        } else {
            // if not... move to node.Righ
            cur = node.Right
        }
	}
	return result
}
```
