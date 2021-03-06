/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。



示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

// func validateStackSequences(pushed []int, popped []int) bool {
// 	indexMap := map[int]int{}
// 	flag := make([]int, len(pushed))

// 	for i, num := range pushed {
// 		indexMap[num] = i
// 	}

// 	lastIndex := -1

// 	for i, num := range popped {
// 		if i == 0 {
// 			lastIndex = indexMap[num]
// 			flag[lastIndex] = 1
// 			continue
// 		}

// 		if indexMap[num] > lastIndex {
// 			lastIndex = indexMap[num]
// 			flag[lastIndex] = 1
// 			continue
// 		}

// 		j := lastIndex - 1
// 		for ; j >= 0; j-- {
// 			if flag[j] == 0 {
// 				break
// 			}
// 		}

// 		if num == pushed[j] {
// 			lastIndex = j
// 			flag[lastIndex] = 1
// 			continue
// 		} else {
// 			return false
// 		}

// 	}

// 	return true
// }

func validateStackSequences(pushed []int, popped []int) bool {
	tmp := make([]int, 0)
	popIndex := 0

	for i, num := range pushed {
		if num == popped[popIndex] {
			popIndex++
			for len(tmp) != 0 && (tmp[len(tmp)-1] == popped[popIndex]) {
				tmp = tmp[0 : len(tmp)-1]
				popIndex++
			}
		} else {
			tmp = append(tmp, num)
		}
	}

	for ; popIndex < len(popped); popIndex++ {
		if popped[popIndex] == tmp[len(tmp)-1] {
			continue
		} else {
			return false
		}
	}

	return true
}
