/*
给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。

另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。

返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）



示例 1：

输入：R = 1, C = 2, r0 = 0, c0 = 0
输出：[[0,0],[0,1]]
解释：从 (r0, c0) 到其他单元格的距离为：[0,1]
*/

package main

type Point struct {
	r int
	c int
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var result [][]int

	flag := make([][]int, R)
	for i := 0; i < R; i++ {
		flag[i] = make([]int, C)
	}

	var queue []Point
	queue = append(queue, Point{r0, c0})

	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]

		if flag[p.r][p.c] == 0 {
			result = append(result, []int{p.r, p.c})
			flag[p.r][p.c] = 1
		}

		if (p.r-1 >= 0) && (flag[p.r-1][p.c] == 0) {
			queue = append(queue, Point{p.r - 1, p.c})
			result = append(result, []int{p.r - 1, p.c})
			flag[p.r-1][p.c] = 1
		}

		if (p.r+1 < R) && (flag[p.r+1][p.c] == 0) {
			queue = append(queue, Point{p.r + 1, p.c})
			result = append(result, []int{p.r + 1, p.c})
			flag[p.r+1][p.c] = 1
		}

		if (p.c-1 >= 0) && (flag[p.r][p.c-1] == 0) {
			queue = append(queue, Point{p.r, p.c - 1})
			result = append(result, []int{p.r, p.c - 1})
			flag[p.r][p.c-1] = 1
		}

		if (p.c+1 < C) && (flag[p.r][p.c+1] == 0) {
			queue = append(queue, Point{p.r, p.c + 1})
			result = append(result, []int{p.r, p.c + 1})
			flag[p.r][p.c+1] = 1
		}
	}

	return result
}
