package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		sum := 0
		for _, no := range queue {
			sum += no.Val
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		if len(lev) == 0 {
			return sum
		}
		queue = lev

	}
	return 0

}
