/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	prev := head

	i := 0

	for {
		if i < n {
			fast = fast.Next

			if fast == nil {
				return head.Next
			}

			i++
			continue
		}

		fast = fast.Next
		prev = slow
		slow = slow.Next

		if fast == nil {
			prev.Next = slow.Next
			slow = nil
			break
		}
	}

	return head
}
