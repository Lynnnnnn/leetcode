/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func generateMatrix(n int) [][]int {
	ele := 1
	i := 0
	j := 0
	curStep := []int{0, 1}

	res := make([][]int, n)

	for i := range res {
		res[i] = make([]int, n)
	}

	for ele <= n*n {
		res[i][j] = ele

		if (i+curStep[0] >= n || i+curStep[0] < 0) ||
			(j+curStep[1] >= n || j+curStep[1] < 0) ||
			(res[i+curStep[0]][j+curStep[1]] != 0) {
			curStep = nextStep(curStep)
		}

		i += curStep[0]
		j += curStep[1]
		ele++
	}

	return res
}
