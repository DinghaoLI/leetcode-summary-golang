# 22. Generate Parentheses

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:
```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

[22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)


```golang
func generateParenthesis(n int) []string {
    if n < 1 { return []string{""} }
    res := []string{}
    // left and right record the number of remaining parenthesis
    var dfs func(left, right int, solution string)
    dfs = func(left, right int, solution string) {
        // finished
        if left == 0 && right == 0 {
            res = append(res, solution)
            return
        }
        // we only have the right parenthesis
        if left == 0 && right > 0 {
            for right>0 {
                solution += ")"
                right--
            }
            res = append(res, solution)
            return
        }
        // if left >= right, we can not add ")"
        if left < right {
            dfs(left, right-1, solution+")")
        }
        // we can add left parenthesis when we have left
        if left > 0 {
            dfs(left-1, right, solution+"(")
        }
    }
    dfs(n, n, "")
    return res
}
```