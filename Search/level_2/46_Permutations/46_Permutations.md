# 46. Permutations

Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
```
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```
[46. Permutations](https://leetcode.com/problems/permutations/)

For each element in a solution we pick different element one by one and record the element that we have already taken to avoid duplication.

```golang
func permute(nums []int) [][]int {
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

        for i := 0; i < length; i++ {
            if !taken[i] {
                // we will use nums[i], so taken[i] == true
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