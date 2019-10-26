# 23. Merge k Sorted Lists

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

```
Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
```

[23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

## Merge one by one

```
Time comlexity: O(l1+l2) + O(l1+l2+l3) + ... + O(l1+l2+ln)
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

 // merge one by one
func mergeKLists(lists []*ListNode) *ListNode {
    // check input
    if len(lists) == 0 { return  nil }
    if len(lists) == 1 { return lists[0] }
    
    // merge one by one
    res := lists[0]
    for i:=1; i<len(lists); i++ {
        res = mergeTwoLists(res, lists[i])
    }
    return res
}

// LC 21. Merge Two Sorted Lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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

## divide and conquer

```
n : length of one list (for each list)
k : number of list
Time complexity: O(nklogk)
Space complexity: O(logk) // depth of stack => O(1)
```

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// divede and conquer
func mergeKLists(lists []*ListNode) *ListNode {
    // check input
    if len(lists) == 0 { return  nil }
    if len(lists) == 1 { return lists[0] }
    
    return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
    // check input
    if left > right { return nil }
    if left == right { return lists[left] }
    if left+1 == right { return mergeTwoLists(lists[left], lists[right]) }
    // recursion for divide and conquer
    mid := (left+right)/2
    l1 := merge(lists, left, mid)
    l2 := merge(lists, mid+1, right)
    return mergeTwoLists(l1, l2)
}

// LC 21. Merge Two Sorted Lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
