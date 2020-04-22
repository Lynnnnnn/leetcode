/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：2
示例 2：

输入：n = 7
输出：21

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func numWays(n int) int {
	if n == 0 {
		return 1
	}

	res := make([]int64, n+1)
	res[0] = 1
	res[1] = 1

	for i := 2; i < n+1; i++ {
		res[i] = res[i-1]%(1e9+7) + res[i-2]%(1e9+7)
	}

	return int(res[n] % (1e9 + 7))
}
