# 542. 01 Matrix

Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.

```

Example 1:

Input:
[[0,0,0],
 [0,1,0],
 [0,0,0]]

Output:
[[0,0,0],
 [0,1,0],
 [0,0,0]]
Example 2:

Input:
[[0,0,0],
 [0,1,0],
 [1,1,1]]

Output:
[[0,0,0],
 [0,1,0],
 [1,2,1]]
```

Note:

The number of elements of the given matrix will not exceed 10,000.
There are at least one 0 in the given matrix.
The cells are adjacent in only four directions: up, down, left and right.

[542. 01 Matrix](https://leetcode.com/problems/01-matrix/)

DP is better than BFS

```golang
// DP 
func updateMatrix(matrix [][]int) [][]int {
    // get the size of matrix
    lr, lc := len(matrix), len(matrix[0])
    MIN := lr*lc
    
    // update matrix[i][j] by matrix[i-1][j] and matrix[i][j-1]
    for i:=0; i<lr; i++ {
        for j:=0; j<lc; j++ {
            if matrix[i][j] == 0 {
                continue
            }
            
            matrix[i][j] = MIN
            
            if i > 0 {
                matrix[i][j] = min(matrix[i][j], matrix[i-1][j]+1)
            }
            if j > 0 {
                matrix[i][j] = min(matrix[i][j], matrix[i][j-1]+1)
            }
        }
    }
    
    // update matrix[i][j] by matrix[i+1][j] and matrix[i][j+1]
    for i:=lr-1; i>=0; i-- {
        for j:=lc-1; j>=0; j-- {
            if matrix[i][j] == 0 {
                continue
            }
            
            if i < lr-1 {
                matrix[i][j] = min(matrix[i][j], matrix[i+1][j]+1)
            }
            if j < lc-1 {
                matrix[i][j] = min(matrix[i][j], matrix[i][j+1]+1)
            }
        }
    }
    return matrix
}

func min(a,b int) int {
    if a < b { return a }
    return b
}
```