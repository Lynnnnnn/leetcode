/*
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。



示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func myPow(x float64, n int) float64 {
	isNegative := false

	if n < 0 {
		isNegative = true
		n *= -1
	}

	result := pow(x, n)

	if isNegative {
		result = 1 / result
	}

	return result
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.00
	}

	result := 1.00

	if n == 1 {
		return x
	}

	if n%2 != 0 {
		tmp := pow(x, (n-1)/2)
		result = tmp * tmp * x
	} else {
		tmp := pow(x, n/2)
		result = tmp * tmp
	}

	return result
}
