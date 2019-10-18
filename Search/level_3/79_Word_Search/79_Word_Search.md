# 79. Word Search

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

```
Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
```

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.

[79. Word Search](https://leetcode.com/problems/word-search/)


```golang
func exist(board [][]byte, word string) bool {
    // check board
    rl := len(board)
    if rl == 0 {
        return false
    }
    cl := len(board[0])
    if cl == 0 {
        return false
    }
    
    // dfs
    var dfs func(idx int, r int, c int) bool
    dfs = func(idx int, r int, c int) bool {
        // find out a solution
        if idx == len(word) { 
            return true 
        }
        
        // check the edge
        if r < 0 || r > rl-1 || c < 0 || c > cl-1 { 
            return false 
        }
        
        // if it's the target letter ?
        if board[r][c] != word[idx] { 
            return false 
        }
        
        // mark board[r][c]
        board[r][c] = '.'
        // dfs
        if dfs(idx+1, r+1, c) || 
            dfs(idx+1, r-1, c) || 
            dfs(idx+1, r, c+1) || 
            dfs(idx+1, r, c-1) {
            return true
        }
        // reset board[r][c]
        board[r][c] = word[idx]
        return false
    }
    
    // dfs from each element
    for i:=0; i<rl; i++ {
        for j:=0; j<cl; j++ {
            if dfs(0, i, j) {
                return true
            }
        }
    }
    return false
}
```