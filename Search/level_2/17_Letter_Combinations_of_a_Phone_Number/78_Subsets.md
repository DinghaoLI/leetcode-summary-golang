# 78. Subsets

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

[78. Subsets](https://leetcode.com/problems/subsets/)

we use the current result to add the next element.

```golang
func subsets(nums []int) [][]int {
    res := [][]int{[]int{}}
    for _,v := range nums {
    	// traverse all the current result
        for _,r := range res {
            // not append(r, []int{v})
            res = append(res, append([]int{v}, r...))
        }
    }
    return res
}
```
