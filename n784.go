/**
给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

示例:
输入: S = "a1b2"
输出: ["a1b2", "a1B2", "A1b2", "A1B2"]
**/

package main

func letterCasePermutation(S string) []string {
	var res []string
	if len(S) == 0 {
		return []string{""}
	}

	leftArray := letterCasePermutation(S[1:])

	for _, v := range leftArray {
		res = append(res, string(S[0])+v)

		if S[0] <= 'Z' && S[0] >= 'A' {
			res = append(res, string(S[0]+'a'-'A')+v)
		} else if S[0] <= 'z' && S[0] >= 'a' {
			res = append(res, string(S[0]-'a'+'A')+v)
		}
	}
	return res
}
