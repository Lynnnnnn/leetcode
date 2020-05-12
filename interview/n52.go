/*
输入两个链表，找出它们的第一个公共节点。
*/

package interview

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	if headA == nil || headB == nil {
		return nil
	}

	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}
