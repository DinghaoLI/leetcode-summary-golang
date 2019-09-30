# 77. Combinations

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

[77. Combinations](https://leetcode.com/problems/combinations/)


```golang
func combine(n int, k int) [][]int {
    res := [][]int{}
    solution := make([]int, k)
    var dfs func(int, int)
    dfs = func(idx, begin int) {
    	// record a result
        if idx == k {
            tmp := make([]int, k)
			copy(tmp, solution)
			res = append(res, tmp)
			return
        }
        // we choose the value for solution[idx]
        // i is bewteen "begin" and n+1-k+idx (the largest value)
        // for example:
        // begin == 1 && idx == 0 && n == 5 && k == 3
        // so, solution[idx] <= 3, because if solution[0] = 4, solution[1] = 5
        // and these is no value for solution[2] 
        for i:=begin; i<=n+1-k+idx; i++ {
            solution[idx] = i
            dfs(idx+1, i+1)
        }
    }
    dfs(0, 1)
    return res
}
```
