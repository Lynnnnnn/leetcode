/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。



示例 1:

输入: [1,2,3,4,5]
输出: True


示例 2:

输入: [0,0,1,2,5]
输出: True

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func isStraight(nums []int) bool {
	min := nums[0]
	max := nums[0]
	change := []int{}
	m := make(map[int]bool)

	i := 0

	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			min = nums[i]
			max = nums[i]
			break
		} else {
			change = append(change, nums[i])
		}
	}

	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			change = append(change, 0)
			continue
		}

		if m[nums[i]] {
			return false
		}

		m[nums[i]] = true

		if nums[i] < min {
			min = nums[i]
			continue
		}

		if nums[i] > max {
			max = nums[i]
			continue
		}
	}

	miss := 0
	for i := min; i <= max; i++ {
		if !m[i] {
			miss++
		}
	}

	if miss <= len(change) {
		return true
	} else {
		return false
	}
}
