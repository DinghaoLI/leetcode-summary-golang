# 303. Range Sum Query - Immutable

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

```
Example:
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
```

Note:
You may assume that the array does not change.
There are many calls to sumRange function.

[303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/)

## DP

```
n: length of nums
m: number of query

Time complexity: O(n) + m * O(1) => O(n+m)
Space complexity: O(n)
```

```golang
type NumArray struct {
    Nums []int
}


func Constructor(nums []int) NumArray {
    // check input
    if len(nums) == 0 {
        return NumArray{ []int{} }
    }
    Nums := make([]int, len(nums))
    // Nums[i]: Nums[0]+...+Nums[i]
    Nums[0] = nums[0]
    for i:=1; i<len(nums); i++ {
        Nums[i] = Nums[i-1]+nums[i] 
    }
    return NumArray{ Nums }
}


func (this *NumArray) SumRange(i int, j int) int {
    if i == 0 { return this.Nums[j] }
    return this.Nums[j] - this.Nums[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
```
