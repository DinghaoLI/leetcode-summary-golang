# 987. Vertical Order Traversal of a Binary Tree

Given a binary tree, return the vertical order traversal of its nodes values.

For each node at position (X, Y), its left and right children respectively will be at positions (X-1, Y-1) and (X+1, Y-1).

Running a vertical line from X = -infinity to X = +infinity, whenever the vertical line touches some nodes, we report the values of the nodes in order from top to bottom (decreasing Y coordinates).

If two nodes have the same position, then the value of the node that is reported first is the value that is smaller.

Return an list of non-empty reports in order of X coordinate.  Every report will have a list of values of nodes.

[987. Vertical Order Traversal of a Binary Tree](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/)

In the article [Go maps in action](https://blog.golang.org/go-maps-in-action)

> When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. Since the release of Go 1.0, the runtime has randomized map iteration order. Programmers had begun to rely on the stable iteration order of early versions of Go, which varied between implementations, leading to portability bugs. If you require a stable iteration order you must maintain a separate data structure that specifies that order. This example uses a separate sorted slice of keys to print a map[int]string in key order:

```golang
import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
```

So, in our solution, we need to use a array to sort the keys in map.

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func verticalTraversal(root *TreeNode) [][]int {
    // store in maps
    m := make(map[int]map[int][]int)
    var dfs func(root *TreeNode, x int, y int)
    dfs = func(root *TreeNode, x int, y int) {
        if root == nil { return }
        if m[x] == nil {
            m[x] = make(map[int][]int)
        }
        if m[x][y] == nil {
            m[x][y] = make([]int, 0)
        } 
        m[x][y] = append(m[x][y], root.Val)
        dfs(root.Left, x-1, y+1)
        dfs(root.Right, x+1, y+1)
    }
    dfs(root, 0, 0)
    
    // sort x keys to ensure the order
    keysX := make([]int, len(m))
    idx := 0
    for k, _ := range m {
        keysX[idx] = k
        idx++
    }
    sort.Ints(keysX)
    res := [][]int{}
    for i:=0; i<len(keysX); i++ {
        // sort y keys to ensure the order
        keysY := make([]int, len(m[keysX[i]]))
        idx := 0
        for k, _ := range m[keysX[i]] {
            keysY[idx] = k
            idx++
        }
        sort.Ints(keysY)

        y := []int{}
        for j:=0; j<len(keysY); j++ {
            sort.Ints(m[keysX[i]][keysY[j]])
            y = append(y, m[keysX[i]][keysY[j]]...)
        }
        res = append(res, y)
    }
    return res
}
```
