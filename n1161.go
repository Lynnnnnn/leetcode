package main

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := root.Val
	minLevel := 1
	up := []*TreeNode{root}
	down := []*TreeNode{}

	level := 1
	for len(up) > 0 {
		sum := 0
		for i := 0; i < len(up); i++ {
			sum += up[i].Val

			if up[i].Left != nil {
				down = append(down, up[i].Left)
			}

			if up[i].Right != nil {
				down = append(down, up[i].Right)
			}
		}

		if sum > max {
			max = sum
			minLevel = level
		} else if sum == max && level < minLevel {
			minLevel = level
		}

		up = down
		down = []*TreeNode{}
		level++
	}

	return minLevel
}
