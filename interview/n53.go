/*
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func search(nums []int, target int) int {
	sum := 0
	middle := 0

	start := 0
	end := len(nums) - 1

	for start <= end {
		middle = (start + end) / 2

		if target == nums[middle] {
			break
		}

		if target > nums[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	for i := middle; i < len(nums); i++ {
		if nums[i] == target {
			sum++
		} else {
			break
		}
	}

	for i := middle - 1; i >= 0; i-- {
		if nums[i] == target {
			sum++
		} else {
			break
		}
	}

	return sum
}
