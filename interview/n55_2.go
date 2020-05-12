/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftBalanced, leftDepth := depth(root.Left)

	if !leftBalanced {
		return false
	}

	rightBalanced, rightDepth := depth(root.Right)

	if !rightBalanced {
		return false
	}

	if math.Abs(float64(rightDepth-leftDepth)) > 1 {
		return false
	} else {
		return true
	}
}

func depth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, left := depth(root.Left)

	if !leftBalanced {
		return false, 0
	}

	rightBalanced, right := depth(root.Right)

	if !rightBalanced {
		return false, 0
	}

	if math.Abs(float64(right-left)) > 1 {
		return false, 0
	}

	if right > left {
		return true, right + 1
	} else {
		return true, left + 1
	}
}
