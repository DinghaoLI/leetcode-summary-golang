# 147. Insertion Sort List

Sort a linked list using insertion sort. 

Algorithm of Insertion Sort:

Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
It repeats until no input elements remain.

```
Example 1:

Input: 4->2->1->3
Output: 1->2->3->4

Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5
```

[147. Insertion Sort List](https://leetcode.com/problems/insertion-sort-list/)

```
Time comlexity: O(n^2)
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

func insertionSortList(head *ListNode) *ListNode {
    headPre := &ListNode{Next: head}

    cur := head
    for cur != nil && cur.Next != nil {
        next := cur.Next
        if cur.Val <= next.Val {
            // in the right order: cur.Val <= next.Val
            cur = next
            continue
        }

        // cur.Val > next.Val: next needs to insert between head and cur
        // delete next here
        cur.Next = next.Next
        
        // insertion
        n1, n2 := headPre, headPre.Next
        // right place for insetion: n1.Val < next.Val <= n2.Val
        for next.Val > n2.Val {
            n1, n2 = n1.Next, n2.Next
        }
        n1.Next, next.Next = next, n2
    }
    return headPre.Next
}
```