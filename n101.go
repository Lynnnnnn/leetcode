/*
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	curLevel := []*TreeNode{root.Left, root.Right}

	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{}
		curLen := len(curLevel)

		for i := 0; i < curLen; i++ {
			if i < curLen/2 {
				if curLevel[i] == nil && curLevel[curLen-i] == nil {
					continue
				}

				if curLevel[i] != nil && curLevel[curLen-1-i] != nil && curLevel[i].Val == curLevel[curLen-1-i].Val {
					nextLevel = append(nextLevel, curLevel[i].Left, curLevel[i].Right)
					continue
				}

				return false
			} else {
				if curLevel[i] != nil {
					nextLevel = append(nextLevel, curLevel[i].Left, curLevel[i].Right)
				}
			}
		}

		curLevel = nextLevel
	}

	return true
}
