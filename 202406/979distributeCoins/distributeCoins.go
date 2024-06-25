package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func distributeCoins(root *TreeNode) int {
	var ans int
	mem := make(map[*TreeNode]pair)

	dfs(root, mem, &ans)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(root *TreeNode, mem1 map[*TreeNode]pair, ans *int) pair {
	if root == nil {
		return pair{}
	}
	if va, ok := mem1[root]; ok {
		return va
	}
	le := dfs(root.Left, mem1, ans)
	ri := dfs(root.Right, mem1, ans)
	pa := pair{
		nodes: le.nodes + ri.nodes + 1,
		vals:  le.vals + ri.vals + root.Val,
	}
	*ans += abs(pa.nodes - pa.vals)
	mem1[root] = pa
	return pa
}

type pair struct {
	nodes int
	vals  int
}
