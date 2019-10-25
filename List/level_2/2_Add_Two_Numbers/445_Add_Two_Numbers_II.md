# 445. Add Two Numbers II

You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

```
Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
```

[445. Add Two Numbers II](https://leetcode.com/problems/add-two-numbers-ii/)


```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // like stack
    s1, s2 := make([]int, 0, 1024), make([]int, 0, 1024)
    for l1 != nil {
        s1 = append(s1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        s2 = append(s2, l2.Val)
        l2 = l2.Next
    }
    
    res := &ListNode{-1, nil}
    cur := res
    carry := 0
    for len(s1) != 0 || len(s2) != 0 || carry != 0 {
        sum := carry
        if len(s1)>0 {
            sum += s1[len(s1)-1]
            s1 = s1[:len(s1)-1]
        }
        if len(s2)>0 {
            sum += s2[len(s2)-1]
            s2 = s2[0:len(s2)-1]
        }
        carry = sum/10
        tmp := &ListNode{sum%10, cur.Next}
        cur.Next = tmp
    }
    return res.Next
}
```
