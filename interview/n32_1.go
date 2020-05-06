/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func levelOrder(root *TreeNode) []int {
	var result []int
	var queue []*TreeNode

	if root == nil {
		return []int{}
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}

		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}

		result = append(result, queue[0].Val)
		queue = queue[1:]
	}

	return result
}
