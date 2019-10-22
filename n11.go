/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// func maxArea(height []int) int {
// 	length := len(height)
// 	max := 0

// 	for i := 0; i < length; i++ {
// 		for j := i; j < length; j++ {
// 			min := min(height[i], height[j])
// 			if min*(j-i) > max {
// 				max = min * (j - i)
// 			}
// 		}
// 	}

// 	return max
// }

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	area := 0

	for i < j {
		if height[i] < height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}

		if area > max {
			max = area
		}
	}

	return max
}
