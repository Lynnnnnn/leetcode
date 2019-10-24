/**
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
**/

package main

func permute(nums []int) [][]int {
	var res [][]int

	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	for i, v := range nums {
		if i == 0 {
			res = append(res, []int{v})
		} else {
			res = newResult(res, v)
		}
	}

	return res
}

func newResult(list [][]int, element int) [][]int {
	var res [][]int

	for _, v := range list {
		tmp := add(v, element)
		res = append(res, tmp...)
	}

	return res
}

func add(list []int, element int) [][]int {
	var res [][]int

	list = append(list, element)

	for i := len(list) - 1; i > 0; i-- {
		newList := make([]int, len(list))
		copy(newList, list)
		res = append(res, newList)

		tmp := list[i-1]
		list[i-1] = list[i]
		list[i] = tmp
	}

	res = append(res, list)

	return res
}
