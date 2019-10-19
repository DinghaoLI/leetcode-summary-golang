# 2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

```
Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

[2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)


```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // check input
    if l1 == nil && l2 == nil { return nil }
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }
    // init 
    res := &ListNode{0, nil}
    cur := res
    carry := 0
    // when one of l1 and l2 is not empty or carry > 0 , => continue
    for l1 != nil || l2 != nil || carry > 0 {
        // init sum
        sum := carry
        // add l1
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        // add l2
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        // carry for the next value
        carry = sum/10
        // creat node for current value and move
        cur.Next = &ListNode{sum%10, nil}
        cur = cur.Next
    }
    return res.Next
}
```
