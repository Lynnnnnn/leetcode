/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)

	if target < 3 {
		return result
	}

	start := 1
	end := 2

	for start < target/2 {
		sum := (start + end) * (end - start + 1) / 2

		if sum == target {
			tmp := make([]int, 0)
			for i := start; i <= end; i++ {
				tmp = append(tmp, i)
			}

			result = append(result, tmp)

			end++
			start++
		} else if sum < target {
			end++
		} else {
			start++
		}
	}

	if target%2 == 1 {
		result = append(result, []int{start, start + 1})
	}

	return result
}
