# 508. Most Frequent Subtree Sum

Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.

Examples 1
Input:
```
  5
 /  \
2   -3
```
return [2, -3, 4], since all the values happen only once, return all of them in any order.
Examples 2
Input:
```
  5
 /  \
2   -5
```
return [2], since 2 happens twice, however -5 only occur once.
Note: You may assume the sum of values in any subtree is in the range of 32-bit signed integer.

[508. Most Frequent Subtree Sum](https://leetcode.com/problems/most-frequent-subtree-sum/)


```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    if root == nil { return []int{} }
    max := math.MinInt32
    res := []int{}
    count := make(map[int]int)
    
    // count the frequence of all the sum of subtree recursively
    var sum func(root *TreeNode) int 
    sum = func(root *TreeNode) int {
        if root == nil { return 0 }
        tmp := root.Val + sum(root.Left) + sum(root.Right)
        count[tmp]++
        return tmp
    }
    sum(root)
    
    // we keep the sums who appear most frequently 
    for k,v := range count {
        if v >= max {
            if v > max {
                max = v
                res = res[0:0]
            }
            res = append(res, k)
        }
    }
    return res
}
```
