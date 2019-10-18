package main

import (
	"fmt"
)

func main() {
	board := [][]byte{[]byte{'a'}}
	fmt.Println(exist(board, "a"))
}

func exist(board [][]byte, word string) bool {
    rl := len(board)
    if rl == 0 {
        return false
    }
    cl := len(board[0])
    if cl == 0 {
        return false
    }

    
    var dfs func(idx int, r int, c int) bool
    dfs = func(idx int, r int, c int) bool {
        fmt.Println(idx)
        if idx == len(word) { 
            return true 
        }
        
        if r < 0 || r > rl-1 || c < 0 || c > cl-1 { 
            return false 
        }
        
        if board[r][c] != word[idx] { 
            return false 
        }
        
        board[r][c] = '.'
        if dfs(idx+1, r+1, c) || 
            dfs(idx+1, r-1, c) || 
            dfs(idx+1, r, c+1) || 
            dfs(idx+1, r, c-1) {
            return true
        }
        board[r][c] = word[idx]
        return false
    }
    
    for i:=0; i<rl-1; i++ {
        for j:=0; j<cl-1; j++ {
        	            	        fmt.Println(i, j)

            if dfs(0, i, j) {

                return true
            }
        }
    }
    return false
}