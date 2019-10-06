# 37. Sudoku Solver

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.

Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.

[37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)


```golang
// DFS
func solveSudoku(board [][]byte) {
    var solve func(k int) bool
    // k is the index of element in board, from 0~80
    solve = func(k int) bool {
        // finish
        if k == 81 {
            return true
        }

        // row and column of this element
        r, c := k/9, k%9
        // board[r][c] is a number => next element
        if board[r][c] != '.' {
            return solve(k + 1)
        }

        // test from 1 to 9
        for b := byte('1'); b <= '9'; b++ {
            if isValid(b, board, r, c) {
                board[r][c] = b
                if solve(k + 1) {
                    return true
                }
            }
        }

        // remove this current element
        board[r][c] = '.'

        return false
    }
    solve(0)
}

func isValid(b byte, board [][]byte, r int, c int) bool {
    // the begin of block for this element
    bi, bj := r/3*3, c/3*3
    for n := 0; n < 9; n++ {
        // verify row, column, block
        if board[r][n] == b ||
            board[n][c] == b ||
            board[bi+n/3][bj+n%3] == b {
            return false
        }
    }
    return true
}
```