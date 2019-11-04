package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// root -> left -> right
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	// init
	stack := []*TreeNode{}
	cur := root

	for len(stack) != 0 || cur != nil {
		// record cur, push cur to stack and turn left (loop)
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		// pop element and go turn right (we have already record the element himself)
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}
	return res
}

// left -> root -> right
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{}
	cur := root
	for len(stack) != 0 || cur != nil{
		// push cur to stack and turn left (loop)
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// pop element, record it and go turn right
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res	
}

// left -> right -> root
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{}
	cur := root
	preVisited := &TreeNode{}
	for len(stack) != 0 || cur != nil{
		// push cur to stack and turn left (loop)
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) != 0 {
			node := stack[len(stack)-1]
			// if <node is leaf> || 
        	// <node have not node.Right> ||
        	// <we have visited node.Right> => we record this node
			if (node.Left == nil && node.Right == nil) || 
			   node.Right == nil || preVisited == node.Right {
				res = append(res, node.Val)
				preVisited = node
				// pop element
				stack = stack[:len(stack)-1]
			} else {
				// turn right
				cur = node.Right
			}
		}
	}
	return res	
}

func main() {
	tree := &TreeNode{1, nil, nil}
	left := &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}
	right :=  &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}
	tree.Left, tree.Right = left, right

	fmt.Println("Preorder Traversal:")
	fmt.Println(preorderTraversal(tree))
	fmt.Println("Inorder Traversal:")
	fmt.Println(inorderTraversal(tree))
	fmt.Println("Postorder Traversal:")
	fmt.Println(postorderTraversal(tree))
}











