/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
*/

package main

func spiralOrder(matrix [][]int) []int {
	var res []int

	row := len(matrix)

	if row == 0 {
		return []int{}
	}

	col := len(matrix[0])
	flag := make([][]int, row)

	for i := range flag {
		flag[i] = make([]int, col)
	}

	i := 0
	j := 0
	count := 0
	curStep := []int{0, 1}

	for count < row*col {
		res = append(res, matrix[i][j])
		flag[i][j] = 1
		count += 1

		if (i+curStep[0] >= row || i+curStep[0] < 0) ||
			(j+curStep[1] >= col || j+curStep[1] < 0) ||
			(flag[i+curStep[0]][j+curStep[1]] == 1) {
			curStep = nextStep(curStep)
		}

		i += curStep[0]
		j += curStep[1]
	}

	return res
}

func nextStep(cur []int) []int {
	if cur[0] == 0 && cur[1] == 1 {
		return []int{1, 0}
	}

	if cur[0] == 1 && cur[1] == 0 {
		return []int{0, -1}
	}

	if cur[0] == 0 && cur[1] == -1 {
		return []int{-1, 0}
	}

	if cur[0] == -1 && cur[1] == 0 {
		return []int{0, 1}
	}

	return []int{0, 0}
}
