/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		middle := (start + end) / 2

		if target == nums[middle] {
			return middle
		}

		if target > nums[middle] {
			if target > nums[end] {
				end -= 1
			} else {
				start += 1
			}
		} else {
			if target < nums[start] {
				start += 1
			} else {
				end -= 1
			}
		}
	}
	return -1
}
