package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 请你反转这棵树中每个 奇数 层的节点值
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	dep := 0
	for len(queue) > 0 {
		dep++
		lev := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Left != nil {
				lev = append(lev, node.Left)
			}
			if node.Right != nil {
				lev = append(lev, node.Right)
			}
		}
		if dep&1 == 1 {
			l, r := 0, len(lev)-1
			for l < r {
				lev[l].Val, lev[r].Val = lev[r].Val, lev[l].Val
				l++
				r--
			}
		}

		queue = lev
	}

	return root
}
