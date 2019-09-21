# 297. Serialize and Deserialize Binary Tree

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Example: 

You may serialize the following tree:

    1
   / \
  2   3
     / \
    4   5

as "[1,2,3,null,null,4,5]"
Clarification: The above format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.s

[297. Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)

leetcode do not support golang, we can only test by ourselves:

```golang
package main

import (
    "fmt"
    "strconv"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func Serialize(root *TreeNode) string {
    if root == nil {
        return "#,"
    }
    str := ""
    str = SerializeTree(root, str)
    return str
}

func SerializeTree(root *TreeNode, str string) string {
    if root == nil {
        str += "#,"
        return str
    }

    str += strconv.Itoa(root.Val) + ","
    str += SerializeTree(root.Left, "")
    str += SerializeTree(root.Right, "")
    return str
}

func Deserialize(str string) *TreeNode {
    if str == "" {
        return nil
    }
    return DeserializeTree(str)
}

func DeserializeTree(str string) *TreeNode {
    index := 0
    var recur func(str string) *TreeNode
    recur = func(str string) *TreeNode {
        if str[index] == '#' {
            index += 2
            return nil
        }
        num := 0
        for str[index] != ',' && index < len(str) {
            num = num*10 + int(str[index] - '0')
            index++
        }
        index++

        root := &TreeNode{num, nil, nil}
        root.Left = recur(str)
        root.Right = recur(str)
        return root
    }
    root := recur(str)
    return root

}

func print(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("%d -> ", root.Val)
    print(root.Left)
    print(root.Right)
}

func main() {
    // q1    1
    // q2  2   3
    // q1 4 5 6 7
    root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}

    fmt.Println("Output: ", Serialize(root))
    fmt.Println("")
    fmt.Println("before")
    print(root)
    fmt.Println("")
    root = Deserialize(Serialize(root))
    fmt.Println("after")
    print(root)
    fmt.Println("")

}
```
