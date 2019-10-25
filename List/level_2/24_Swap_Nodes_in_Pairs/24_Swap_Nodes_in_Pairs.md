# 24. Swap Nodes in Pairs

Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

```

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.

```

[24. Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)


```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    // check input
    if head == nil || head.Next == nil {
        return head
    }
    // init
    res := &ListNode{-1, head}
    parent, cur, forward := res, head, head.Next
    
    for forward != nil {
        // swap
        parent.Next = forward
        cur.Next = forward.Next
        forward.Next = cur
        
        // check the next pair
        if cur == nil || cur.Next == nil || cur.Next.Next == nil {
            break
        }
    
        // move to next pair
        parent, cur, forward = cur, cur.Next, cur.Next.Next
    }
    return res.Next
}
```
