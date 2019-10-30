# 268. Missing Number

Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

```
Example 1:

Input: [3,0,1]
Output: 2

Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
```

Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

[268. Missing Number](https://leetcode.com/problems/missing-number/)


```golang
// using sum to calculate the missing value  
func missingNumber(nums []int) int {
    n := len(nums)
    sum := (0+n)*(n+1)/2
    for _,v := range nums {
        sum -= v
    }
    return sum
}

// XOR 
// using 0^0=0, a^0=a, b^b=0, a^b^b=a
// the missing one will be xor only one time, but others will be xor twice
func missingNumber(nums []int) int {
    x:=0
    for i:=1; i<=len(nums); i++ {
        x = x^i^nums[i-1]
    }
    return x
}
```
