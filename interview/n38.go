/*
输入一个字符串，打印出该字符串中字符的所有排列。



你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。



示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	if len(s) == 1 {
		return []string{s}
	}

	exist := make(map[string]bool)

	res := make([]string, 0)

	for i, c := range s {
		newS := s[0:i] + s[i+1:]
		newArr := permutation(newS)

		if len(newArr) == 0 {
			if !exist[string(c)] {
				res = append(res, string(c))
				exist[string(c)] = true
			}
		} else {
			for _, item := range newArr {
				if !exist[string(c)+item] {
					res = append(res, string(c)+item)
					exist[string(c)+item] = true
				}
			}
		}
	}

	return res
}
