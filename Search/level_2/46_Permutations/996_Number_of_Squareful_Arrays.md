# 996. Number of Squareful Arrays

Given an array A of non-negative integers, the array is squareful if for every pair of adjacent elements, their sum is a perfect square.

Return the number of permutations of A that are squareful.  Two permutations A1 and A2 differ if and only if there is some index i such that A1[i] != A2[i].


Example 1:

Input: [1,17,8]
Output: 2
Explanation: 
[1,8,17] and [17,8,1] are the valid permutations.
Example 2:

Input: [2,2,2]
Output: 1

Note:

1 <= A.length <= 12
0 <= A[i] <= 1e9

[996. Number of Squareful Arrays](https://leetcode.com/problems/number-of-squareful-arrays/)

Based on the ``LC 47. Permutation II``, 

```golang
func numSquarefulPerms(A []int) int {
	length := len(A)
	count := 0
	// taken[i] == true means A[i] is picked by solution
	taken := make([]bool, length)
	solution := make([]int, length)

	var dfs func(idx int)
	dfs = func(idx int) {
		// find out a solution
		if idx == len(A) {
			count++
			return
		}

		// in the same depth, we don't want to use the same value, so we record the value
		// that we have used
		used := make(map[int]bool, length-idx)

		for i := 0; i < length; i++ {
			if !taken[i] && !used[A[i]] {

				// make sure that this element+previous element can be
				// a perfect square => backtracking
				if idx > 0 && !isPerfectSquare(solution[idx-1]+A[i]) {
					continue
				}

				// we will use A[i], so taken[i] = true and used[A[i]] = true
				used[A[i]] = true
				taken[i] = true
				solution[idx] = A[i]

				dfs(idx + 1)

				// in the next iteration, we pick A[i+1]
				taken[i] = false
			}
		}
	}

	dfs(0)

	return count
}

// LC 367. Valid Perfect Square
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 0, num/2
	for left <= right {
		mid := (left + right) / 2
		midValue := mid * mid
		if midValue == num {
			return true
		} else if midValue < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
```