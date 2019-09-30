# 40. Combination Sum II

Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]

[40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)


```golang
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}
    
    var dfs func (candidates []int, target int)
    dfs = func (candidates []int, target int) {
        if target == 0 {
            // find out a solution
            tmp := make([]int, len(solution))
            copy(tmp, solution)
		    res = append(res, tmp)
	    }

	    if len(candidates) == 0 || target < candidates[0] {
		    // target < candidates[0] 因为candidates是排序好的
		    return
	    }

	    // use candidates[0]
        solution = append(solution, candidates[0])
	    dfs(candidates[1:], target-candidates[0])
        solution = solution[:len(solution)-1]
        
	    // do not use candidates[0] and those duplicate values
	    dfs(next(candidates), target)
    }
	dfs(candidates, target)

	return res
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}
```
