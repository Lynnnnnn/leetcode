/**
当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 A 的最大湍流子数组的长度

示例 1：

输入：[9,4,2,10,7,8,8,1,9]
输出：5
解释：(A[1] > A[2] < A[3] > A[4] < A[5])
*/
package main

func maxTurbulenceSize(A []int) int {
	length := len(A)

	if length == 1 {
		return length
	}

	start := 0
	end := 1
	for end < length-1 {
		if A[end] == A[start] {
			start++
			end++
		} else {
			break
		}
	}

	max := end - start + 1
	if end >= length-1 {
		return 1
	}

	for end < length-1 {
		if check(A[end-1], A[end], A[end+1]) {
			end = end + 1
			if end-start+1 > max {
				max = end - start + 1
			}
		} else {
			start = end
			end = start + 1
		}
	}

	return max
}

func check(a int, b int, c int) bool {
	if (a < b && b > c) || (a > b && b < c) {
		return true
	} else {
		return false
	}
}
