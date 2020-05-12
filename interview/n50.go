/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

//map
// func firstUniqChar(s string) byte {
// 	m := make(map[rune]int)

// 	for _, c := range s {
// 		if _, ok := m[c]; ok {
// 			m[c] += 1
// 		} else {
// 			m[c] = 1
// 		}
// 	}

// 	for _, c := range s {
// 		if m[c] == 1 {
// 			return byte(c)
// 		}
// 	}

// 	return ' '
// }

//array
func firstUniqChar(s string) byte {
	arr := make([]int, 27)

	for _, c := range s {
		arr[c-'a'] += 1
	}

	for _, c := range s {
		if arr[c-'a'] == 1 {
			return byte(c)
		}
	}

	return ' '
}
