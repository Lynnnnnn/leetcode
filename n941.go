/**
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

A.length >= 3
在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[B.length - 1]


示例 1：

输入：[2,1]
输出：false
示例 2：

输入：[3,5,5]
输出：false
示例 3：

输入：[0,3,2,1]
输出：true
*/
package main

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	leftEndFlag := false
	rightEndFlag := false

	start := 0
	end := len(A) - 1

	for start < end {
		if A[start] < A[start+1] {
			start += 1
		} else {
			leftEndFlag = true
		}

		if A[end] < A[end-1] {
			end -= 1
		} else {
			rightEndFlag = true
		}

		if leftEndFlag && rightEndFlag {
			break
		}
	}

	if start == end && start != 0 && start != len(A)-1 {
		return true
	} else {
		return false
	}
}
