/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if (root == p) || (root == q) {
		return root
	}

	var left *TreeNode = nil
	var right *TreeNode = nil

	if root.Left != nil {
		left = lowestCommonAncestor(root.Left, p, q)
	}

	if root.Right != nil {
		right = lowestCommonAncestor(root.Right, p, q)
	}

	if (root.Left == nil) && (root.Right == nil) {
		return nil
	}

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}
