package main

import (
	"sort"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	ans := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		th := 0
		for _, no := range queue {
			// 这样写更不容易出错
			th += no.Val
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		ans = append(ans, th)
		queue = lev
	}
	sort.Ints(ans)
	if len(ans) < k {
		return -1
	}
	return int64(ans[len(ans)-k])
}

func kthLargestLevelSum1(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	ans := []int{root.Val}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		th := 0
		for _, no := range queue {
			if no.Left != nil {
				lev = append(lev, no.Left)
				th += no.Left.Val
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
				th += no.Right.Val
			}
		}
		if len(lev) > 0 {
			ans = append(ans, th)
		}

		queue = lev
	}
	sort.Ints(ans)
	if len(ans) < k {
		return -1
	}
	return int64(ans[len(ans)-k])
}
