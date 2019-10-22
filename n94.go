/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var left []int
	var right []int

	if root.Left != nil {
		left = inorderTraversal(root.Left)
	}

	if root.Right != nil {
		right = inorderTraversal(root.Right)
	}

	left = append(left, root.Val)

	if len(right) > 0 {
		left = append(left, right...)
	}
	return left
}
