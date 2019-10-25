# 206. Reverse Linked List

Reverse a singly linked list.

```
Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
```

[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

## Recursively

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// recursively
func reverseList(head *ListNode) *ListNode {
    // check input
    if head == nil || head.Next == nil {
        return head
    }
    // reverse from head.Next, and get: head -> newTail <- ... <-newHead
    // set it to: nil <- head <- newTail <- ... <-newHead
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}

```


## Iteratively

```golang
// iteratively
func reverseList(head *ListNode) *ListNode {
    // check input
    if head == nil || head.Next == nil {
        return head
    }
    // init the first one
    res := &ListNode{-1, head}
    first := head
    first.Next, head = nil, head.Next
    // head insertion for reversing list
    for head != nil {
        tmp := res.Next
        res.Next, head = head, head.Next
        res.Next.Next = tmp
    }
    return res.Next
}
```
