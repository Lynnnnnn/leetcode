/*
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。



你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。



示例 1:

输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func twoSum(n int) []float64 {
	if n == 1 {
		return []float64{float64(1) / 6, float64(1) / 6, float64(1) / 6, float64(1) / 6, float64(1) / 6, float64(1) / 6}
	}

	last := twoSum(n - 1)
	result := make([]float64, 5*n+1)

	for i := 0; i < len(last); i++ {
		result[i] = float64(1)/6 * last[i]
	}

	for i := 1; i < 6; i++ {
		for j := 0; j < len(last); j++ {
			result[i+j] += float64(1) / 6 * last[j]
		}
	}

	return result
}
