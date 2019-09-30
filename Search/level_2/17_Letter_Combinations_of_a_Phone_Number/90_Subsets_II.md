# 90. Subsets II

Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

[90. Subsets II](https://leetcode.com/problems/subsets-ii/)


```golang
func subsetsWithDup(nums []int) [][]int {
    // sorted
    sort.Ints(nums)

    res := [][]int{}
    solution := []int{}
    // backtrace
    var dfs func(idx int)
    dfs = func(idx int) {
        tmp := make([]int, len(solution))
        copy(tmp, solution)
        res = append(res, tmp)
        for i:=idx; i<len(nums); i++ {
            // remove duplicates elements
            if i > idx && nums[i] == nums[i-1] {
                continue
            }
            // use nums[i]
            solution = append(solution, nums[i])
            dfs(i+1)
            // clean slice for the next value
            solution = solution[:len(solution)-1]
        }
    }
    
    dfs(0)
    return res
}
```
