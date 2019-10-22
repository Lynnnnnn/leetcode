package main

import "fmt"

func main() {
	// fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// 	// result := validSquare([]int{0, 0}, []int{2, 2}, []int{2, 0}, []int{0, 2})
	// 	// result := validSquare([]int{0, 0}, []int{0, 0}, []int{0, 0}, []int{0, 0})
	// result := totalFruit([]int{3, 3, 4})
	// 	// fmt.Println(result)
	// result := rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3})
	// result := coinChange([]int{186, 419, 83, 408}, 6249)
	// result := coinChange([]int{83, 186}, 1265)
	// result := longestConsecutive([]int{1, 2, 0, 1})
	// result := maxProfit([]int{7, 6, 4, 3, 1})
	// result := longestPalindrome("cbbd")
	// result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	// result := uniquePaths(7, 3)
	// result := subsets([]int{9, 0, 3, 5, 7})
	// result := grayCode(0)
	// result := findKthLargest([]int{0, 1, 2, 3, 4, 18, 10, 11, 0}, 1)
	result := productExceptSelf([]int{4, 3, 2, 1, 2})
	fmt.Println(result)
}
