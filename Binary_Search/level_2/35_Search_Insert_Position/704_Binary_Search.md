# 704. Binary Search


Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.

```
Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
 

Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].
```

[704. Binary Search](https://leetcode.com/problems/binary-search/)

```
二分搜索模版:

- 边界条件 left <= right
- mid = (left+right)/2, 根据比较条件移动 left用left = mid+1， right用right = mid-1
- 最后返回 left

注意以上实现方式有一个好处，就是当循环结束时，如果没有找到目标元素，那么left一定停在恰好比目标大的index上，right一定停在恰好比目标小的index上，所以个人比较推荐这种实现方式。
```

```
Time complexity: O(logn)
Space comlexity: O(1)
```

## binary search

```golang
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left+right)/2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}
```
