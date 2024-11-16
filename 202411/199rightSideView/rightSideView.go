package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	q := []*TreeNode{root}
	for len(q) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range q {
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		if len(lev) == 0 {
			break
		}
		ans = append(ans, lev[len(lev)-1].Val)
		q = lev
	}
	return ans
}
