# 17. Letter Combinations of a Phone Number

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.


Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

[17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)

```
Input: "23"
'2': []string{"a", "b", "c"},
'3': []string{"d", "e", "f"},

Initialize:

res => {""}

For 2:

res => {""+"a", ""+"b", ""+"c"}

For 3:

res => {"a"+"d", "a"+"e", "a"+"f", "b"+"d",..., "c"+"f"}

```

## DFS

Time Complexity: O(4^n)

Space Complexity: O(4^n + n)

```golang
var m = map[byte][]byte{
    '2': []byte{'a', 'b', 'c'},
    '3': []byte{'d', 'e', 'f'},
    '4': []byte{'g', 'h', 'i'},
    '5': []byte{'j', 'k', 'l'},
    '6': []byte{'m', 'n', 'o'},
    '7': []byte{'p', 'q', 'r', 's'},
    '8': []byte{'t', 'u', 'v'},
    '9': []byte{'w', 'x', 'y', 'z'},
}

// dfs
func letterCombinations(digits string) []string {
    if digits == "" { return []string{} }
    list := []byte(digits)
    
    res := []string{}
    solution := []byte{}
    var dfs func(idx int)
    dfs = func(idx int) {
        // get a solution
        if idx == len(list) {
            res = append(res, string(solution[:]))
            return
        }
        for i:=0; i<len(m[list[idx]]); i++ {
            // add value
            solution = append(solution, m[list[idx]][i])
            dfs(idx+1)
            // clean
            solution = solution[:len(solution)-1]
        }
    }
    dfs(0)
    return res
}
```

## BFS

Time Complexity: O(4^n)

Space Complexity: O(4^n x 2)

```golang
// the map of number => characters
var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}


// bfs + combination => the next combination is based on the previous result
func letterCombinations(digits string) []string {
	// digits is "" => empty slice
    if digits == "" { return []string{} }
    list := []byte(digits)
    
    res := []string{""}
    // for each number
    for i:=0; i<len(list); i++ {
        tmp := []string{}
        // for each character
        for j:=0; j<len(m[list[i]]); j++ {
        	// add each character to the end of all the previous result
            for k:=0; k<len(res); k++ {
                tmp = append(tmp, res[k]+m[list[i]][j])
            }
        }
        res = tmp
    }
    return res
}
```
