/*
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。



示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

//到叶子节点的路径才算
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return [][]int{[]int{sum}}
	}

	left := pathSum(root.Left, sum-root.Val)
	right := pathSum(root.Right, sum-root.Val)

	if len(left) == 0 && len(right) == 0 {
		return [][]int{}
	}

	merged := append(left, right...)

	result := make([][]int, 0)

	for _, arr := range merged {
		tmp := append([]int{root.Val}, arr...)
		result = append(result, tmp)
	}

	return result
}
