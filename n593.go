/**
给定二维空间中四点的坐标，返回四点是否可以构造一个正方形。

一个点的坐标（x，y）由一个有两个整数的整数数组表示。

示例:

输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
输出: True


注意:

所有输入整数都在 [-10000，10000] 范围内。
一个有效的正方形有四个等长的正长和四个等角（90度角）。
输入点没有顺序。
*/
package main

import "math"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p12 := distancePow2(p1, p2)
	p13 := distancePow2(p1, p3)
	p23 := distancePow2(p2, p3)

	if p12 == p13 && p12 != 0 {
		if p12+p13 != p23 {
			return false
		}

		p24 := distancePow2(p2, p4)
		p34 := distancePow2(p3, p4)

		if p24 == p34 && p24+p34 == p23 {
			return true
		} else {
			return false
		}
	} else {
		var maxPoint []int
		var maxDistance float64

		if p12 > p13 {
			maxPoint = p2
			maxDistance = p12
		} else {
			maxPoint = p3
			maxDistance = p13
		}

		p14 := distancePow2(p1, p4)
		pm4 := distancePow2(maxPoint, p4)

		if p14 == pm4 && p14 != 0 && p14+pm4 == maxDistance {
			return true
		} else {
			return false
		}
	}
}

func distancePow2(p1 []int, p2 []int) float64 {
	return math.Pow(math.Abs(float64(p1[0]-p2[0])), 2) + math.Pow(math.Abs(float64(p1[1]-p2[1])), 2)
}
