/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func longestPalindrome(s string) string {
	length := len(s)

	dp := make([][]int, length)

	for i := range dp {
		dp[i] = make([]int, length)
	}

	longest := ""
	max := 0

	for n := 0; n < length; n++ {
		i := 0
		j := n

		for {
			if i >= length || j >= length {
				break
			}

			if (i + 1) <= (j - 1) {
				if s[i] == s[j] && dp[i+1][j-1] == 1 {
					dp[i][j] = 1

					if j-i+1 > max {
						max = j - i + 1
						longest = s[i : j+1]
					}
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = 1

					if j-i+1 > max {
						max = j - i + 1
						longest = s[i : j+1]
					}
				}
			}

			i++
			j++
		}
	}

	return longest
}
