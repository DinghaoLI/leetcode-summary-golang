# 746. Min Cost Climbing Stairs

On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

```
Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.

Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].

Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
```

[746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)

## DP

``
Time complexity: O(n)
Space complexity: O(n)
``

```golang
func minCostClimbingStairs(cost []int) int {
    // check input
    if len(cost) == 0 { return 0 }
    if len(cost) == 1 { return cost[0] }
    if len(cost) == 2 { return min(cost[0], cost[1]) }
    
    // dp[i] means Min Cost Climbing Stairs at step i, the last step has no cost
    dp := make([]int, len(cost)+1)
    copy(dp, cost)
    for i:=2; i<len(dp); i++ {
        // trasition: dp[i] cost dp[i] + min(dp[i-1], dp[i-2])
        dp[i] = dp[i] + min(dp[i-1], dp[i-2])
    }
    return dp[len(dp)-1]
}

func min(a,b int) int {
    if a < b { return a }
    return b
}
```

## DP space complexity O(1)

``
Time complexity: O(n)
Space complexity: O(1)
``

```golang
func minCostClimbingStairs(cost []int) int {
    // check input
    if len(cost) == 0 { return 0 }
    if len(cost) == 1 { return cost[0] }
    if len(cost) == 2 { return min(cost[0], cost[1]) }
    
    // dp[i] means Min Cost Climbing Stairs at step i, the last step has no cost
    dp1, dp2 := 0, 0
    for i:=2; i<len(cost)+1; i++ {
        // trasition
        dp := min(dp1+cost[i-1], dp2+cost[i-2])
        dp1, dp2 = dp, dp1
    }
    return dp1
}

func min(a,b int) int {
    if a < b { return a }
    return b
}
```

