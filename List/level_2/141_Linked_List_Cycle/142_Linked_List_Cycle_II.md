# 142. Linked List Cycle II

Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

``` 
Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.


Example 2:

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.


Example 3:

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.


 

Follow-up:
Can you solve it without using extra space?
```

[142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)


```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
    // check input
    if head == nil || head.Next == nil {
        return nil
    }
    
    // slow takes 1 step and fast takes 2 step for each iteration  
    slow, fast := head, head
    cycle := false
    for fast.Next != nil && fast.Next.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
        if fast == slow { 
            cycle = true
            break 
        }
    }
    // no cycle
    if !cycle { return  nil }
    // n = length from head to entrance
    // m = length of cycle
    // k = length of from entrance to slow 
    // m-k = the length from meeting point to the next entrance
    // slow = n + k
    // fast = qm + n + k  (q>0, number of turns)
    // fast = 2 * slow
    // => n = qm -k
    // => n = (q-1)m + (m-k)
    // so we start from slow and head, and they are moving forward at the same speed，When they meet, the meeting point is the entrance of cycle
    detector := head
    for detector != slow {
        detector, slow = detector.Next, slow.Next
    }
    
    return detector
}
```
