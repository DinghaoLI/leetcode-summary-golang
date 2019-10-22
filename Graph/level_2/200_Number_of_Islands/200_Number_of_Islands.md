# 200. Number of Islands

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

```
Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3
```

[200. Number of Islands](https://leetcode.com/problems/number-of-islands/)

## DFS

Time Complexity : O(mn)
Space Complexity : O(mn)


```golang
// DFS
func numIslands(grid [][]byte) int {
    // check out
    lr := len(grid)
    if lr == 0 { return 0 }
    lc := len(grid[0])
    if lc == 0 { return 0 }
    
    // dfs to clean the connected island
    var dfs func(r,c int) bool
    dfs = func(r,c int) bool {
        if r < 0 || r >= lr || c < 0 || c >= lc {
            return false
        }
        if grid[r][c] == '0' {
            return false
        }
        grid[r][c] = '0'
        dfs(r+1,c)
        dfs(r-1,c)
        dfs(r,c+1)
        dfs(r,c-1)
        return true
    }
    
    // iteration 
    count := 0
    for r:=0; r<lr; r++ {
        for c:=0; c<lc; c++ {
            if dfs(r, c) {
                count++
            }
        }
    }
    return count
}
```