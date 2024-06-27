package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Left == nil && no.Right == nil {
				ans += no.Val
			}
			if no.Left != nil {
				no.Left.Val = no.Val*2 + no.Left.Val
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				no.Right.Val = no.Val*2 + no.Right.Val
				lev = append(lev, no.Right)
			}
		}
		queue = lev
	}
	return ans
}
