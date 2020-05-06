/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package interview

func verifyPostorder(postorder []int) bool {
	length := len(postorder)

	if length == 0 {
		return true
	}

	last := postorder[len(postorder)-1]
	i := 0

	for ; i < len(postorder)-1; i++ {
		if postorder[i] > last {
			break
		}
	}

	for j := i; j < len(postorder)-1; j++ {
		if postorder[j] < last {
			return false
		}
	}

	if i == len(postorder)-2 {
		return true
	}

	return verifyPostorder(postorder[0:i]) && verifyPostorder(postorder[i:length-1])
}
