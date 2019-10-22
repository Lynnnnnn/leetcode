/*
格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。

给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。

示例 1:

输入: 2
输出: [0,1,3,2]
解释:
00 - 0
01 - 1
11 - 3
10 - 2

对于给定的 n，其格雷编码序列并不唯一。
例如，[0,2,3,1] 也是一个有效的格雷编码序列。

00 - 0
10 - 2
11 - 3
01 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gray-code
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	nextGrayCode := grayCode(n - 1)

	for i := len(nextGrayCode) - 1; i >= 0; i-- {
		val := nextGrayCode[i]
		nextGrayCode = append(nextGrayCode, (1<<(uint)(n-1) + val))
	}

	return nextGrayCode
}
