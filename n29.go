/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
**/

package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	isPositive := true

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		isPositive = false
	}

	if dividend < 0 {
		dividend *= -1
	}

	if divisor < 0 {
		divisor *= -1
	}

	res := 0

	for dividend >= divisor {
		dividend -= divisor
		res--
	}

	if isPositive {
		if float64(res*(-1)) > (math.Pow(2, 31) - 1) {
			return int(math.Pow(2, 31) - 1)
		} else {
			return res * (-1)
		}
	} else {
		return res
	}
}
