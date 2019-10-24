/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	res := ""
	min := len(strs[0])

	for _, v := range strs {
		if len(v) < min {
			min = len(v)
		}
	}

	for i := 0; i < min; i++ {
		c := strs[0][i]
		for _, v := range strs {
			if v[i] != c {
				return res
			}
		}

		res += string(c)
	}

	return res
}
