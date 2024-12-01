package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// bfs 做法
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	q := []*TreeNode{root}
	fa := make(map[*TreeNode]*TreeNode) // 父节点
	dfs1(root, nil, fa)
	ans := make([]int, 0)
	visit := make(map[*TreeNode]bool)
	q = []*TreeNode{target}
	visit[target] = true
	dis := 0

	for len(q) > 0 {
		if dis == k {
			for _, no := range q {
				ans = append(ans, no.Val)
			}
			return ans
		}
		lev := make([]*TreeNode, 0)
		for _, no := range q {
			nex := []*TreeNode{no.Left, no.Right, fa[no]}
			for _, nx := range nex {
				if nx != nil && !visit[nx] {
					visit[nx] = true
					lev = append(lev, nx)
				}
			}
		}
		q = lev
		dis++
	}

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
