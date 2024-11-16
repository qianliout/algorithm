package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	ans = append(ans, []int{root.Val})
	for len(q) > 0 {
		lev1 := make([]*TreeNode, 0)
		lev2 := make([]int, 0)
		for _, no := range q {
			if no.Left != nil {
				lev1 = append(lev1, no.Left)
				lev2 = append(lev2, no.Left.Val)
			}
			if no.Right != nil {
				lev1 = append(lev1, no.Right)
				lev2 = append(lev2, no.Right.Val)
			}
		}
		if len(lev2) == 0 {
			break
		}
		ans = append(ans, lev2)
		q = lev1
	}
	return ans
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	flag := false
	q := []*TreeNode{root}
	ans = append(ans, []int{root.Val})
	for len(q) > 0 {
		lev1 := make([]*TreeNode, 0)
		lev2 := make([]int, 0)
		for _, no := range q {
			if no.Left != nil {
				lev1 = append(lev1, no.Left)
				lev2 = append(lev2, no.Left.Val)
			}
			if no.Right != nil {
				lev1 = append(lev1, no.Right)
				lev2 = append(lev2, no.Right.Val)
			}
		}
		if len(lev2) == 0 {
			break
		}
		if flag {
			ans = append(ans, lev2)
			flag = !flag
		} else if !flag {
			l, r := 0, len(lev2)-1
			for l < r {
				lev2[l], lev2[r] = lev2[r], lev2[l]
				l++
				r--
			}

			ans = append(ans, lev2)
			flag = !flag
		}

		q = lev1
	}
	return ans
}
