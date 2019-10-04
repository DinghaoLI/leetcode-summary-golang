# 784. Letter Case Permutation

Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.

Examples:
Input: S = "a1b2"
Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

Input: S = "3z4"
Output: ["3z4", "3Z4"]

Input: S = "12345"
Output: ["12345"]

Note:

S will be a string with length between 1 and 12.
S will consist only of letters or digits.

[784. Letter Case Permutation](https://leetcode.com/problems/letter-case-permutation/)


```golang
func letterCasePermutation(S string) []string {
    res := []string{}
    list := []byte(S)
    var dfs func(depth int)
    dfs = func(depth int) {
        if depth == len(list) {
            res = append(res, string(list))
            return
        }
        if !isLetter(list[depth]) {
            // element is not letter => skip
            dfs(depth+1)
            return
        } else {
            // pick letter
            dfs(depth+1)

            //change case and pick it
            list[depth] = changeCase(list[depth])
            dfs(depth+1)

            // recover
            list[depth] = changeCase(list[depth])
        }
    }
    dfs(0)
    return res
}


func isLetter(l byte) bool {
    if l>='a' && l<='z' || l>='A' && l<='Z' {
        return true
    }
    return false
}

func changeCase(l byte) byte {
    if l>='a' && l<='z'{
        return l-32
    }
    return l+32
}
```