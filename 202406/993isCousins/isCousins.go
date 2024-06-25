package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// bfs 不好搞
func isCousins(root *TreeNode, x int, y int) bool {
	a := dfs(root, nil, x, 0)
	b := dfs(root, nil, y, 0)
	if a != nil && b != nil && a.pa != b.pa && a.dep == b.dep {
		return true
	}

	return false
}

func dfs(root, pa *TreeNode, x int, dep int) *pire {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return &pire{dep: dep, pa: pa}
	}
	le := dfs(root.Left, root, x, dep+1)
	ri := dfs(root.Right, root, x, dep+1)
	if le != nil {
		return le
	}
	if root != nil {
		return ri
	}
	return nil
}

type pire struct {
	dep int
	pa  *TreeNode
}
