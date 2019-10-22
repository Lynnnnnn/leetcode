/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TODO 未完成
package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 || len(nums) == 0 {
		return 0
	}

	pivot := nums[0]
	start := 1
	end := len(nums) - 1

	for start < end {
		for nums[end] >= pivot && start < end {
			end--
		}

		if start < end {
			nums[start] = nums[end]
			start++
		}

		for nums[start] <= pivot && start < end {
			start++
		}

		if start < end {
			nums[end] = nums[start]
			end--
		}
	}

	nums[start] = pivot

	findKthLargest(nums[0:start], 1)
	findKthLargest(nums[start+1:], 1)
	fmt.Println(nums)

	return 0
}
