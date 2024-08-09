package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func amountOfTime(root *TreeNode, start int) int {
	var startNode *TreeNode
	faM := make(map[*TreeNode]*TreeNode)
	var dfs1 func(fa *TreeNode)

	dfs1 = func(fa *TreeNode) {
		if fa == nil {
			return
		}
		if fa.Val == start {
			startNode = fa
		}
		if fa.Left != nil {
			faM[fa.Left] = fa
			dfs1(fa.Left)
		}
		if fa.Right != nil {
			faM[fa.Right] = fa
			dfs1(fa.Right)
		}
	}
	dfs1(root)
	// 每个节点的值 互不相同
	visit := make(map[int]bool)
	visit[start] = true
	ans := 0

	queue := []*TreeNode{startNode}

	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			fa := faM[no]
			le := no.Left
			ri := no.Right
			if fa != nil && !visit[fa.Val] {
				visit[fa.Val] = true
				lev = append(lev, fa)
			}
			if ri != nil && !visit[ri.Val] {
				visit[ri.Val] = true
				lev = append(lev, ri)
			}
			if le != nil && !visit[le.Val] {
				visit[le.Val] = true
				lev = append(lev, le)
			}
		}
		if len(lev) > 0 {
			ans++
		}
		queue = lev
	}
	return ans
}
