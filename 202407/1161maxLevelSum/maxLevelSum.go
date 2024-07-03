package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mx := root.Val
	res := 1
	dep := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		dep++
		sum := 0
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			sum += no.Val
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}

		queue = lev
		if sum > mx {
			mx = max(mx, sum)
			res = dep
		}
	}
	return res
}
