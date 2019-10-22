/*
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/product-of-array-except-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func productExceptSelf(nums []int) []int {
	length := len(nums)

	if length == 2 {
		return []int{nums[1], nums[0]}
	}

	res := make([]int, length)
	copy(res, nums)

	i := 1
	j := len(nums) - 2
	middle := len(nums) / 2

	for i < middle {
		res[i] *= res[i-1]
		i++
	}

	for j > middle {
		res[j] *= res[j+1]
		j--
	}

	var tmpj int = res[middle-1]
	var tmpi int = res[middle] * res[middle+1]

	for j := middle; j < length; j++ {
		if j == length-1 {
			res[j] = tmpj
		} else {
			res[j] = tmpj * res[j+1]
			tmpj = tmpj * nums[j]
		}
	}

	for i := middle - 1; i >= 0; i-- {
		if i == 0 {
			res[i] = tmpi
		} else {
			res[i] = res[i-1] * tmpi
			tmpi = tmpi * nums[i]
		}
	}

	return res
}
