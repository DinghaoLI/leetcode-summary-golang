# 33. Search in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

```
Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

[33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

```
n: the length of nums
Time complexity: O(logn)
Space comlexity: O(1)
```
```
二分搜索模版:

- 边界条件 left <= right
- 移动 left用left = mid+1， right用right = mid-1
- 最后返回 left

注意判断单调区间时需要用<=和>=
```


```golang
func search(nums []int, target int) int {
    if len(nums) == 0 { return -1 }
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left+right)/2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[left] { // left~mid is increasing
            if nums[left] <= target && target <= nums[mid] { // target is in left~mid
                right = mid-1
            } else { // target is in mid~right
                left = mid+1
            }
        } else { // mid~right is increasing
            if nums[mid] <= target && target <= nums[right] { // target is in mid~right
                left = mid+1
            } else { // target is in left~mid
                right = mid-1
            }
        }
    }
    return -1
}
```