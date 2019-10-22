# 934. Shortest Bridge

In a given 2D binary array A, there are two islands.  (An island is a 4-directionally connected group of 1s not connected to any other 1s.)

Now, we may change 0s to 1s so as to connect the two islands together to form 1 island.

Return the smallest number of 0s that must be flipped.  (It is guaranteed that the answer is at least 1.)

```
Example 1:
Input: [[0,1],[1,0]]
Output: 1

Example 2:
Input: [[0,1,0],[0,0,0],[0,0,1]]
Output: 2

Example 3:
Input: [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
Output: 1
 

Note:

1 <= A.length = A[0].length <= 100
A[i][j] == 0 or A[i][j] == 1
```

[934. Shortest Bridge](https://leetcode.com/problems/shortest-bridge/)


## DFS + BFS

- Use DFS to find one island and color all the nodes as 2.
- Use BFS to find the shortest path from any nodes with color 2 to any nodes with color 1.

Time complexity: O(mn)
Space complexity: O(mn)


```golang
// pair of row and column
type Pair struct {
    R int
    C int
}

func shortestBridge(A [][]int) int {
     // check out
    lr := len(A)
    if lr == 0 { return -1 }
    lc := len(A[0])
    if lc == 0 { return -1 }
    
    // bfs queue
    q := []Pair{}
    // dfs to find the first island
    found := false
    for r:=0; r<lr && !found; r++ {
        for c:=0; c<lc && !found; c++ {
            if A[r][c] == 1 {
                // mark the first island 1 => 2
                dfs(A, r, c, &q)
                found = true
            }
        }
    }
    
    // use bfs to expand the first island until another island is reached
    steps := 0
    // trick: to simulate {r+1, c}, {r-1, c}, {r, c+1}, {r, c-1}
    dirs := []int{0, 1, 0, -1, 0}
    for len(q) != 0 {
        for size:=len(q); size>0; size-- {
            r, c := q[0].R, q[0].C
            q = q[1:]
            // iterate {r+1, c}, {r-1, c}, {r, c+1}, {r, c-1}
            for i:=0; i<4; i++ {
                tr, tc := r+dirs[i], c+dirs[i+1]
                if tr<0 || tr>=lr || tc<0 || tc>=lc || A[tr][tc] == 2 {
                    continue
                }
                // another island is reached
                if A[tr][tc] == 1 { return steps }
                // take a step
                A[tr][tc] = 2
                q = append(q, Pair{tr, tc})
            }
        }
        steps++
    }
    return steps
}

// dfs like LC 200. Number of Islands
func dfs(A [][]int, r int, c int, q *[]Pair) {
    if r < 0 || r >= len(A) || c < 0 || c >= len(A[0]) {
        return 
    }
    if A[r][c] == 0 || A[r][c] == 2 {
        return 
    }
    *q = append(*q, Pair{r, c})
    A[r][c] = 2
    dfs(A, r+1,c, q)
    dfs(A, r-1,c, q)
    dfs(A, r,c+1, q)
    dfs(A, r,c-1, q)
}
```




