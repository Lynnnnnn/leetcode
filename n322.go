/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
**/

package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		min := amount + 1
		for _, v := range coins {
			if i-v >= 0 {
				if dp[i-v] < min {
					min = dp[i-v]
				}
			}
		}
		dp[i] = min + 1
	}

	if dp[len(dp)-1] >= amount+1 {
		return -1
	} else {
		return dp[len(dp)-1]
	}
}
