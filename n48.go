/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func rotate(matrix [][]int) {
	length := len(matrix)
	middle := (length) / 2

	for n := 0; n < middle; n++ {
		for i := n; i < length-1-n; i++ {
			indexI := n
			indexJ := i

			tmp := matrix[indexI][indexJ]

			for step := 0; step < 4; step++ {
				nextI := indexJ
				nextJ := length - 1 - indexI

				tmp2 := matrix[nextI][nextJ]
				matrix[nextI][nextJ] = tmp
				tmp = tmp2

				indexI = nextI
				indexJ = nextJ
			}
		}
	}
}
