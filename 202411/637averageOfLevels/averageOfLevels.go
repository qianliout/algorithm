package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, float64(root.Val))
	q := []*TreeNode{root}
	for len(q) > 0 {
		lev := make([]*TreeNode, 0)
		sum := 0
		for _, no := range q {
			if no.Left != nil {
				lev = append(lev, no.Left)
				sum += no.Left.Val
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
				sum += no.Right.Val
			}
		}
		if len(lev) == 0 {
			break
		}
		ans = append(ans, float64(sum)/float64(len(lev)))
		q = lev
	}
	return ans
}
