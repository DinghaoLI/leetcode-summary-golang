# 148. Sort List

Sort a linked list in O(n log n) time using constant space complexity.

```
Example 1:

Input: 4->2->1->3
Output: 1->2->3->4

Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5
```

[148. Sort List](https://leetcode.com/problems/sort-list/)

- to array => sort => to list 
- divide and conquer => merge sort

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// divide and conquer => merge sort (up to bottom)
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }
    // split
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    mid := slow.Next
    slow.Next = nil
    // merge
    return merge(sortList(head), sortList(mid))
}

func merge(h1, h2 *ListNode) *ListNode {
    head := &ListNode{0, nil}
    tmp := head
    for h1 != nil && h2 != nil {
        if h1.Val > h2.Val {
            h1, h2 = h2, h1
        }
        tmp.Next = h1
        h1 = h1.Next
        tmp = tmp.Next
    }
    if h1 != nil {
        tmp.Next = h1
    }
    if h2 != nil {
        tmp.Next = h2
    }
    return head.Next
}
```
