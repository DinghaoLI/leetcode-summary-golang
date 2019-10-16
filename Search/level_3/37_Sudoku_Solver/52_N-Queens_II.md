# 52. N-Queens II

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

```
Example:

Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
```

[52. N-Queens II](https://leetcode.com/problems/n-queens-ii/)


```golang
func totalNQueens(n int) int {
    if n == 0 {
        return 0
    }

    // record the usage of '+' 
    cols := make([]bool, n)
    // record the usage of '\' 
    d1 := make([]bool, 2*n-1)
    // record the usage of '/' 
    d2 := make([]bool, 2*n-1)

    board := make([]string, n)

    res := 0

    var dfs func(r int)
    dfs = func(r int) {

        if r == len(board) {
            res++
            return
        }

        n := len(board)

        for c := 0; c < len(board); c++ {
            // calculate the index 
            id1 := n + (r-c) - 1
            id2 := 2*n - (r+c) - 2
            if !cols[c] && !d1[id1] && !d2[id2] {
                b := make([]byte, n)
                for i := range b {
                    b[i] = '.'
                }
                b[c] = 'Q'
                board[r] = string(b)
                // set index
                cols[c], d1[id1], d2[id2] = true, true, true

                dfs(r + 1)

                // reset index
                cols[c], d1[id1], d2[id2] = false, false, false
            }
        }
    }
    dfs(0)

    return res
}
```