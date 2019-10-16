/**
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
**/

//TODO 使用map
package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v] = 1
	}

	var uniqNums []int

	for k, _ := range numMap {
		uniqNums = append(uniqNums, k)
	}

	length := len(uniqNums)

	if length <= 1 {
		return length
	}

	sort.Ints(uniqNums)

	max := 1
	start := 0
	end := 1

	for end < length {
		if uniqNums[end]-uniqNums[end-1] == 1 {
			if end-start+1 > max {
				max = end - start + 1
			}

			end += 1
		} else {
			start = end
			end += 1
		}
	}

	return max
}
