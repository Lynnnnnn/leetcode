/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？



示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}

// 	max := 0

// 	i := 0

// 	for ; i < len(prices)-1; i++ {
// 		if prices[i] < prices[i+1] {
// 			break
// 		}
// 	}

// 	if i == len(prices)-1 {
// 		return 0
// 	}

// 	min := prices[i]

// 	for i = i + 1; i < len(prices); i++ {
// 		if prices[i]-min > max {
// 			max = prices[i] - min
// 			continue
// 		}

// 		if prices[i] < min {
// 			min = prices[i]
// 		}
// 	}

// 	return max
// }

func maxProfitDP(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	dp[0] = 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			dp[i] = dp[i-1]
			continue
		}

		profit := prices[i] - min

		if profit > dp[i-1] {
			dp[i] = profit
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(prices)-1]
}
