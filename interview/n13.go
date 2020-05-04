/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}

	flag := make([][]int, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]int, n)
	}

	return 1 + dfs(0, 0, m, n, flag, k)
}

func dfs(row, col, m, n int, flag [][]int, k int) int {
	flag[row][col] = 1

	if !checkK(row, col, k) {
		return 0
	}

	res := 0
	if row-1 >= 0 && flag[row-1][col] == 0 && checkK(row-1, col, k) {
		res += 1 + dfs(row-1, col, m, n, flag, k)
	}

	if col+1 < n && flag[row][col+1] == 0 && checkK(row, col+1, k) {
		res += 1 + dfs(row, col+1, m, n, flag, k)
	}

	if row+1 < m && flag[row+1][col] == 0 && checkK(row+1, col, k) {
		res += 1 + dfs(row+1, col, m, n, flag, k)
	}

	if col-1 >= 0 && flag[row][col-1] == 0 && checkK(row, col-1, k) {
		res += 1 + dfs(row, col-1, m, n, flag, k)
	}

	return res
}

func checkK(row, col, k int) bool {

	if row/10+row%10+col/10+col%10 > k {
		return false
	}

	return true
}
