/*
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func levelOrder2(root *TreeNode) [][]int {
	result := make([][]int, 0)
	resultNode := make([][]*TreeNode, 0)

	if root == nil {
		return result
	}

	resultNode = append(resultNode, []*TreeNode{root})

	for {
		tmp := make([]*TreeNode, 0)
		last := resultNode[len(resultNode)-1]

		for _, item := range last {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}

			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}

		if len(tmp) == 0 {
			break
		}

		resultNode = append(resultNode, tmp)
	}

	for _, arr := range resultNode {
		tmp := make([]int, 0)

		for _, node := range arr {
			tmp = append(tmp, node.Val)
		}

		result = append(result, tmp)
	}

	return result
}
