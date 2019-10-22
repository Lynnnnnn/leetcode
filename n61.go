/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func rotateRight(head *ListNode, k int) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	fast := head
// 	slow := head

// 	for i := 0; i < k; i++ {
// 		if fast.Next == nil {
// 			fast = head
// 		} else {
// 			fast = fast.Next
// 		}
// 	}

// 	for fast.Next != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}

// 	fast.Next = head
// 	head = slow.Next
// 	slow.Next = nil

// 	return head
// }

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	var arr []*ListNode

	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	length := len(arr)
	k = k % length

	if k == 0 {
		return arr[0]
	} else {
		arr[length-1].Next = arr[0]
		index := length - k
		arr[index-1].Next = nil
		return arr[index]
	}
}
