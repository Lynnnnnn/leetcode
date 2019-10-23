/*
在一个给定的数组nums中，总是存在一个最大元素 。

查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

如果是，则返回最大元素的索引，否则返回-1。

示例 1:

输入: nums = [3, 6, 1, 0]
输出: 1
解释: 6是最大的整数, 对于数组中的其他整数,
6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.


示例 2:

输入: nums = [1, 2, 3, 4]
输出: -1
解释: 4没有超过3的两倍大, 所以我们返回 -1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func dominantIndex(nums []int) int {
	if len(nums) < 1 {
		return -1
	}

	if len(nums) == 1 {
		return 0
	}

	maxindex := 0
	lastmaxindex := 0

	if nums[0] > nums[1] {
		maxindex = 0
		lastmaxindex = 1
	} else {
		maxindex = 1
		lastmaxindex = 0
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[maxindex] {
			lastmaxindex = maxindex
			maxindex = i
		} else if nums[i] > nums[lastmaxindex] {
			lastmaxindex = i
		}
	}

	if nums[maxindex] >= 2*nums[lastmaxindex] {
		return maxindex
	}

	return -1
}
