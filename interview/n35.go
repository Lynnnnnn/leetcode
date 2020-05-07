/*
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var copyList, tmpCopy *Node
	var copyAddr []*Node
	addrOrder := make(map[*Node]int)

	i := 0
	tmpHead := head

	for tmpHead != nil {
		addrOrder[tmpHead] = i
		i++

		if copyList == nil {
			copyList = &Node{tmpHead.Val, tmpHead.Next, nil}
			tmpCopy = copyList

			copyAddr = append(copyAddr, copyList)
		}

		if tmpHead.Next != nil {
			tmpCopy.Next = &Node{tmpHead.Next.Val, tmpHead.Next.Next, nil}
			copyAddr = append(copyAddr, tmpCopy.Next)
		}

		tmpHead = tmpHead.Next
		tmpCopy = tmpCopy.Next
	}

	tmpHead = head
	tmpCopy = copyList

	for tmpHead != nil {
		if tmpHead.Random != nil {
			idx := addrOrder[tmpHead.Random]
			tmpCopy.Random = copyAddr[idx]
		}

		tmpHead = tmpHead.Next
		tmpCopy = tmpCopy.Next
	}

	return copyList
}

//TODO 间隔法牛逼
