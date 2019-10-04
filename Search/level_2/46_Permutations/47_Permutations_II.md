# 47. Permutations II

Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:
```
Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```
[47. Permutations II](https://leetcode.com/problems/permutations-ii/submissions/)

For each element in a solution we pick different element one by one and record the element that we have already taken and the value that we have already used to avoid duplication.

```golang
func permuteUnique(nums []int) [][]int {
    length := len(nums)
    solution := make([]int, length)
    // taken[i] == true means nums[i] is picked by solution
    taken := make([]bool, length)

    var res [][]int
    
    var dfs func(idx int)
    dfs = func(idx int) {
        // find out a solution
        if idx == len(nums) {
            tmp := make([]int, length)
            copy(tmp, solution)
            res = append(res, tmp)
            return
        }

        // in the same depth, we don't want to use the same value, so we record the value 
        // that we have used 
        used := make(map[int]bool, length-idx)
        
        for i := 0; i < length; i++ {
            if !taken[i] && !used[nums[i]]  {
                
                // we will use nums[i], so taken[i] = true and used[nums[i]] = true
                used[nums[i]] = true
                taken[i] = true
                solution[idx] = nums[i]

                dfs(idx+1)

                // in the next iteration, we pick nums[i+1]
                taken[i] = false
            }
        }
    }

    dfs(0)

    return res
}
```