# 70. Climbing Stairs

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

```
Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

[70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

## DP

``
Time complexity: O(n)
Space complexity: O(n)
``

```golang
func climbStairs(n int) int {
    // check input
    if n < 3 { return n }
    
    // dp[i-1] means that we have dp[i-1] ways to reach step i   
    dp := make([]int, n)
    
    dp[0], dp[1] = 1, 2
    for i:=2; i<n; i++ {
        // trasition: to reach step i, we can start from i-1 or i-2
        dp[i] = dp[i-1]+ dp[i-2]
    }
    return dp[n-1]
}
```

## DP space complexity O(1)

``
Time complexity: O(n)
Space complexity: O(1)
``

```golang
func climbStairs(n int) int {
    // check input
    if n < 3 { return n }    
    
    // for dp[i], we only need dp[i-1], dp[i-2]
    res, dp0, dp1 := 0, 1, 2
    for i:=2; i<n; i++ {
        // trasition: to reach step i, we can start from i-1 or i-2
        res = dp0 + dp1
        dp0, dp1 = dp1, res
    }
    return res
}
```

