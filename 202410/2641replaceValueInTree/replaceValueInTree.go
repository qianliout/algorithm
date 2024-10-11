package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Val = 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]*TreeNode, 0)
		childSum := 0
		for _, no := range queue {
			if no.Left != nil {
				level = append(level, no.Left)
				childSum += no.Left.Val
			}
			if no.Right != nil {
				level = append(level, no.Right)
				childSum += no.Right.Val
			}
		}
		for _, no := range queue {
			s := 0
			if no.Left != nil {
				s += no.Left.Val
			}
			if no.Right != nil {
				s += no.Right.Val
			}
			if no.Left != nil {
				no.Left.Val = childSum - s
			}
			if no.Right != nil {
				no.Right.Val = childSum - s
			}
		}
		queue = level
	}
	return root
}
