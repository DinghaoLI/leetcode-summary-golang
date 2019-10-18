# 212. Word Search II

Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

```
Example:

Input: 
board = [
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
words = ["oath","pea","eat","rain"]

Output: ["eat","oath"]
```

[212. Word Search II](https://leetcode.com/problems/word-search-ii/)

## DFS + Trie

```
l : length of word
Time complexity: O(m*n*4^l)
Space complexity: O(l)
```

```golang
// DFS + Trie

// TrieNode structure
type TrieNode struct {
    next [26]*TrieNode
    word string
}

// build TrieNode from words
func buildTrie(words []string) *TrieNode {
    root := new(TrieNode)
    for _, word := range words {
        cur := root
        for _, c := range word {
            cidx := int(c - 'a')
            if cur.next[cidx] == nil {
                cur.next[cidx] = new(TrieNode)
            }
            cur = cur.next[cidx]
        }
        cur.word = word
    }

    return root
}


func findWords(board [][]byte, words []string) []string {
    res := []string{}

    // check board
    rl := len(board)
    if rl == 0 {
        return res
    }
    cl := len(board[0])
    if cl == 0 {
        return res
    }
        
    // dfs
    var dfs func(r int, c int, trie *TrieNode) 
    dfs = func(r int, c int, trie *TrieNode) {
         // check the edge
        if r < 0 || r > rl-1 || c < 0 || c > cl-1 { 
            return 
        }
        
        // check trie
        letter := board[r][c]
        if letter == '.' || trie.next[int(letter-'a')] == nil {
            return
        }
    
        // move to the next node/letter
        trie = trie.next[int(letter-'a')]
        
        // find out a solution
        if len(trie.word) > 0 { 
            res = append(res, trie.word)
            trie.word = ""
        }
        
        // mark board[r][c]
        board[r][c] = '.'
        // dfs
        dfs(r+1, c, trie) 
        dfs(r-1, c, trie)
        dfs(r, c+1, trie) 
        dfs(r, c-1, trie)
        
        // reset board[r][c]
        board[r][c] = letter
    }
    
    // creat trie
    trie := buildTrie(words)
    
    // dfs from each element
    for i:=0; i<rl; i++ {
        for j:=0; j<cl; j++ {
            dfs(i, j, trie)
        }
    }
    return res
}
```

## DFS

```
l : length of word
Time complexity: O(sum(m*n*4^l))
Space complexity: O(l)
```

```golang
// DFS
// based on LC 79. Word Search
// complexity = len(words) * complexity(LC 79. Word Search)
func findWords(board [][]byte, words []string) []string {
    res := []string{}
    // for each word
    for _, str := range words {
        if exist(board, str) {
            res = append(res, str)
        }
    }
    return res
}

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
        res := dfs(idx+1, r+1, c) || 
            dfs(idx+1, r-1, c) || 
            dfs(idx+1, r, c+1) || 
            dfs(idx+1, r, c-1)
        
        // reset board[r][c]
        board[r][c] = word[idx]
        return res
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

