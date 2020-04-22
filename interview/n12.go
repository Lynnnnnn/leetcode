/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				flag := make([][]int, row)
				for k := 0; k < row; i++ {
					flag[k] = make([]int, col)
				}

				exist := exitWithFlag(board, word, i, j, row, col, flag)

				if exist {
					return true
				}
			}
		}
	}

	return false
}

func exitWithFlag(board [][]byte, word string, startRow, startCol, row, col int, flag [][]int) bool {
	if len(word) == 1 {
		return true
	}

	flag[startRow][startCol] = 1

	if (startRow+1) < row && board[startRow+1][startCol] == word[1] && flag[startRow+1][startCol] == 0 {
		if exitWithFlag(board, word[1:], startRow+1, startCol, row, col, flag) {
			return true
		}
	}

	if startRow-1 >= 0 && board[startRow-1][startCol] == word[1] && flag[startRow-1][startCol] == 0 {
		if exitWithFlag(board, word[1:], startRow-1, startCol, row, col, flag) {
			return true
		}
	}

	if startCol+1 < col && board[startRow][startCol+1] == word[1] && flag[startRow][startCol+1] == 0 {
		if exitWithFlag(board, word[1:], startRow, startCol+1, row, col, flag) {
			return true
		}
	}

	if startCol-1 >= 0 && board[startRow][startCol-1] == word[1] && flag[startRow][startCol-1] == 0 {
		if exitWithFlag(board, word[1:], startRow, startCol-1, row, col, flag) {
			return true
		}
	}

	flag[startRow][startCol] = 0

	return false
}
