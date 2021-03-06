/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。



示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

// func exchange(nums []int) []int {
// 	var n1 []int
// 	var n2 []int

// 	for _, n := range nums {
// 		if n%2 == 0 {
// 			n2 = append(n2, n)
// 		} else {
// 			n1 = append(n1, n)
// 		}
// 	}

// 	n1 = append(n1, n2...)

// 	return n1
// }

//首尾指针
func exchange(nums []int) []int {
	start := 0
	end := len(nums) - 1

	for start < end {
		for start < len(nums) && nums[start]%2 == 1 {
			start++
		}

		for end >= 0 && nums[end]%2 == 0 {
			end--
		}

		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}

	return nums
}
