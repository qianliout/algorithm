package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// dfs 做法
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	fa := make(map[*TreeNode]*TreeNode) // 父节点
	dfs1(root, nil, fa)
	ans := make([]int, 0)
	visit := make(map[*TreeNode]bool)
	var dfs func(node *TreeNode, k int)
	dfs = func(node *TreeNode, k int) {
		if node == nil {
			return
		}
		if visit[node] {
			return
		}
		visit[node] = true
		if k == 0 {
			ans = append(ans, node.Val)
		}
		next := []*TreeNode{node.Left, node.Right, fa[node]}
		for _, nex := range next {
			dfs(nex, k-1)
		}
	}
	dfs(target, k)
	return ans
}

func dfs1(node, fa *TreeNode, faM map[*TreeNode]*TreeNode) {
	faM[node] = fa
	if node.Left != nil {

		dfs1(node.Left, node, faM)
	}
	if node.Right != nil {
		dfs1(node.Right, node, faM)
	}
}
