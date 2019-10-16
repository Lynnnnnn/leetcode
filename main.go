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
	result := countBinarySubstrings("00110011")
	fmt.Println(result)
}
