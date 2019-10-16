/*
给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。

示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
*/

package main

// func countBinarySubstrings(s string) int {
// 	var countArray []int

// 	total := 0
// 	count := 1

// 	for i := 1; i < len(s); i++ {
// 		if s[i] == s[i-1] {
// 			count++
// 		} else {
// 			countArray = append(countArray, count)
// 			count = 1
// 		}
// 	}

// 	countArray = append(countArray, count)

// 	for i := 1; i < len(countArray); i++ {
// 		if countArray[i] <= countArray[i-1] {
// 			total += countArray[i]
// 		} else {
// 			total += countArray[i-1]
// 		}
// 	}

// 	return total
// }

func countBinarySubstrings(s string) int {
	total := 0
	prev := 0
	cur := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cur++
		} else {
			prev = cur
			cur = 1
		}

		if prev >= cur {
			total++
		}
	}

	return total
}
