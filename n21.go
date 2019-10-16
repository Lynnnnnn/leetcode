/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
**/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cursor *ListNode

	for l1 != nil && l2 != nil {
		node := new(ListNode)
		node.Next = nil

		if l1.Val <= l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}

		if cursor == nil {
			cursor = node
			head = node
		} else {
			cursor.Next = node
			cursor = cursor.Next
		}
	}

	if l1 == nil && l2 == nil {
		return head
	}

	var leftl *ListNode
	if l1 != nil {
		leftl = l1
	} else {
		leftl = l2
	}

	for leftl != nil {
		node := ListNode{leftl.Val, nil}

		if cursor == nil {
			cursor = &node
			head = cursor
		} else {
			cursor.Next = &node
			cursor = cursor.Next
		}

		leftl = leftl.Next
	}

	return head
}
