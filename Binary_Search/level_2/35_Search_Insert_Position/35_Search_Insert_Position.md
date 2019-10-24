# 35. Search Insert Position


Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

```
Example 1:

Input: [1,3,5,6], 5
Output: 2

Example 2:

Input: [1,3,5,6], 2
Output: 1

Example 3:

Input: [1,3,5,6], 7
Output: 4

Example 4:

Input: [1,3,5,6], 0
Output: 0
```

[35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)

```
二分搜索模版:

- 边界条件 left <= right
- 移动 left用left = mid+1， right用right = mid-1
- 最后返回 left

注意以上实现方式有一个好处，就是当循环结束时，如果没有找到目标元素，那么left一定停在恰好比目标大的index上，right一定停在恰好比目标小的index上，所以个人比较推荐这种实现方式。
```

```
Time complexity: O(logn)
Space comlexity: O(1)
```

## binary search - upper bound

```golang
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left+right)/2
        if target == nums[mid] {
            return mid
        }
        if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return left
}
```
