# 39. Combination Sum

Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

[39. Combination Sum](https://leetcode.com/problems/combination-sum/)


```golang
func combinationSum(candidates []int, target int) [][]int {
	// sort small -> big
    sort.Ints(candidates)
    res := [][]int{}
    tmp := []int{}
    var dfs func(start int, target int)
    dfs = func(start int, target int) {
        for i:=start ; i< len(candidates); i++ {
            if candidates[i] == target {
            	// get a combination
                tmp = append(tmp, candidates[i])
                ans := make([]int, len(tmp))
                copy(ans, tmp)
                res = append(res, ans)
                tmp = tmp[:len(tmp)-1]
                break
            } else if candidates[i] > target {
            	// candidates is sorted, so we don't need to test the next/bigger element
                break
            } else {
            	// find the next element (can use itself)
                tmp = append(tmp, candidates[i])
                dfs(i, target-candidates[i])
                tmp = tmp[:len(tmp)-1]
            }
        }
    }
    dfs(0, target)
    return res
}
```
