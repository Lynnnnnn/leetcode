/*
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

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
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var levelNodes []*TreeNode

	if root != nil {
		levelNodes = []*TreeNode{root}
		res = append(res, []int{root.Val})
	}

	for {
		var currents []int
		var tmpLevelNodes []*TreeNode

		for _, p := range levelNodes {
			if p.Left != nil {
				currents = append(currents, p.Left.Val)
				tmpLevelNodes = append(tmpLevelNodes, p.Left)
			}

			if p.Right != nil {
				currents = append(currents, p.Right.Val)
				tmpLevelNodes = append(tmpLevelNodes, p.Right)
			}
		}

		if len(currents) != 0 {
			res = append(res, currents)
			levelNodes = tmpLevelNodes
		} else {
			break
		}
	}

	return res
}
