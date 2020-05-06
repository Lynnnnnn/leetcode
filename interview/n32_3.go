/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。


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
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func levelOrder3(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	resultNode := make([][]*TreeNode, 0)
	resultNode = append(resultNode, []*TreeNode{root})

	level := 1

	for {
		tmp := make([]*TreeNode, 0)
		last := resultNode[len(resultNode)-1]

		for i := len(last) - 1; i >= 0; i-- {
			if level%2 == 1 {
				if last[i].Right != nil {
					tmp = append(tmp, last[i].Right)
				}

				if last[i].Left != nil {
					tmp = append(tmp, last[i].Left)
				}
			} else {
				if last[i].Left != nil {
					tmp = append(tmp, last[i].Left)
				}

				if last[i].Right != nil {
					tmp = append(tmp, last[i].Right)
				}
			}
		}

		if len(tmp) == 0 {
			break
		}

		resultNode = append(resultNode, tmp)
		level++
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
