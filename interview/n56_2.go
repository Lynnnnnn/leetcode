/*
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。



示例 1：

输入：nums = [3,4,3,3]
输出：4
示例 2：

输入：nums = [9,1,7,9,7,9,7]
输出：1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		if count, ok := m[n]; ok {
			m[n] = count + 1
		} else {
			m[n] = 1
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}
