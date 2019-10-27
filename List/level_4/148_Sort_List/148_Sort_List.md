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

## Merge Sort (bottom up)

```
Merge Sort (bottom up):

log(n) rounds, in the i-th round, we split the list into n/(2^i) lists (n, n/2, n/4), each group has 2^i elements.

merge every pair of sorted lists 

Time comlexity: O(nlogn)
Space comlexity: O(1)
```

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// bottom up
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }
    
    // calculate the length of list
    length := 0
    for cur := head; cur != nil; {
        length++
        cur = cur.Next
    }
    
    preHead := &ListNode{-1, head}
    // logn times iterations
    for n:=1; n<length; n <<= 1 {
        cur := preHead.Next
        tail := preHead
        // one complet iteration of bottom-up merge
        for cur != nil {
            left := cur
            right := split(left, n)
            cur = split(right, n)
            tail.Next, tail = mergeTwoLists(left, right)
        }
    }
    
    return preHead.Next
}

// split the list in two parts: fitst n element and rest
// return the head of the rest
func split(head *ListNode, n int) *ListNode {
    if head == nil { return nil }
    for n > 1 && head != nil {
        head = head.Next
        n--
    }
    if head == nil { return nil }
    target := head.Next 
    head.Next = nil
    return target
}

// LC 21. Merge Two Sorted Lists + return head, tail
func mergeTwoLists(l1, l2 *ListNode) (*ListNode,*ListNode) {
    // init result
    res := &ListNode{-1, nil}
    cur := res
    
    // add the smaller one to the result list
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next, l1 = l1, l1.Next
            cur.Next.Next = nil
            cur = cur.Next
        } else {
            cur.Next, l2 = l2, l2.Next
            cur.Next.Next = nil
            cur = cur.Next
        }
    }
    
    // add the rest list
    if l1 != nil {
        cur.Next = l1
    }
     if l2 != nil {
        cur.Next = l2
    }
    
    if res.Next == nil { return nil, nil }
    head, tail := res.Next, res.Next.Next
    for tail != nil && tail.Next != nil {
        tail = tail.Next
    }
    return head, tail
}
```

## Merge Sort (top down)

```
Merge Sort (top down):

l1, l2 = split(l)
l1', l2' = sortList(l1), sortList(l2)
merge(l1', l2')

Time comlexity: O(nlogn)
Space comlexity: O(logn)
```

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// bottom up
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
    return mergeTwoLists(sortList(head), sortList(mid))
}

// LC 21. Merge Two Sorted Lists
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    // init result
    res := &ListNode{-1, nil}
    cur := res
    
    // add the smaller one to the result list
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next, l1 = l1, l1.Next
            cur.Next.Next = nil
            cur = cur.Next
        } else {
            cur.Next, l2 = l2, l2.Next
            cur.Next.Next = nil
            cur = cur.Next
        }
    }
    
    // add the rest list
    if l1 != nil {
        cur.Next = l1
    }
     if l2 != nil {
        cur.Next = l2
    }
    return res.Next
}
```