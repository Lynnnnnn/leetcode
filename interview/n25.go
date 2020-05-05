/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, tmp *ListNode

	for l1 != nil || l2 != nil {
		var val int

		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				val = l1.Val
				l1 = l1.Next
			} else {
				val = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l2.Val
			l2 = l2.Next
		}

		if result == nil {
			result = &ListNode{
				Val:  val,
				Next: nil,
			}

			tmp = result
		} else {
			tmp.Next = &ListNode{
				Val:  val,
				Next: nil,
			}

			tmp = tmp.Next
		}
	}

	return result
}
