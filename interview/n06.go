/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/

package interview

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append([]int{head.Val}, result...)
		head = head.Next
	}

	return result
}
