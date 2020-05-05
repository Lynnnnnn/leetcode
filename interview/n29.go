/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

import "fmt"

type Step struct {
	Row, Col int
}

func (s *Step) NextStep() *Step {
	if s.Row == 0 && s.Col == 1 {
		return &Step{1, 0}
	}

	if s.Row == 1 && s.Col == 0 {
		return &Step{0, -1}
	}

	if s.Row == 0 && s.Col == -1 {
		return &Step{-1, 0}
	}

	if s.Row == -1 && s.Col == 0 {
		return &Step{0, 1}
	}

	return nil
}

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)

	row := len(matrix)

	if row == 0 {
		return result
	}

	col := len(matrix[0])

	flag := make([][]int, row)

	for i := 0; i < row; i++ {
		flag[i] = make([]int, col)
	}

	s := &Step{0, 1}
	iterRow := 0
	iterCol := 0

	for i := 0; i < row*col; i++ {
		flag[iterRow][iterCol] = 1
		result = append(result, matrix[iterRow][iterCol])
		fmt.Println(result)

		if iterRow+s.Row < row && iterRow+s.Row >= 0 && iterCol+s.Col < col && iterCol+s.Col >= 0 && flag[iterRow+s.Row][iterCol+s.Col] == 0 {
			iterRow += s.Row
			iterCol += s.Col
		} else {
			s = s.NextStep()
			iterRow += s.Row
			iterCol += s.Col
		}
	}

	return result
}
