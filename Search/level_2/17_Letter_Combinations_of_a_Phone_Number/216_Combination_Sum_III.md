# 216. Combination Sum III

Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

Note:

All numbers will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: k = 3, n = 7
Output: [[1,2,4]]
Example 2:

Input: k = 3, n = 9
Output: [[1,2,6], [1,3,5], [2,3,4]]

[216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)


```golang
func combinationSum3(k int, n int) [][]int {
    res := [][]int{}
    solution := []int{}
    var dfs func(target int, idx int)
    dfs = func(target int, idx int) {
        // find a solution
        if len(solution) == k && target == 0 {
            tmp := make([]int, len(solution))
            copy(tmp, solution)
            res = append(res, tmp)
            return
        }
        // test others
        for i:=idx; i<=9 && len(solution) <= k; i++ {
            solution = append(solution, i)
            dfs(target-i, i+1)
            solution = solution[:len(solution)-1]
        }
    }
    dfs(n, 1)
    return res
}
```
