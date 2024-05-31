package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func pruneTree(root *TreeNode) *TreeNode {
	mem := make(map[*TreeNode]bool)
	return dfs(root, mem)
}

func dfs(root *TreeNode, mem map[*TreeNode]bool) *TreeNode {
	if root == nil {
		return nil
	}
	left := contain(root.Left, mem)
	if !left {
		root.Left = nil
	}
	right := contain(root.Right, mem)
	if !right {
		root.Right = nil
	}

	if !right && !left && root.Val != 1 {
		return nil
	}

	dfs(root.Left, mem)
	dfs(root.Right, mem)
	return root
}

func contain(root *TreeNode, mem map[*TreeNode]bool) bool {
	if root == nil {
		return false
	}
	if va, ok := mem[root]; ok {
		return va
	}
	if root.Val == 1 {
		return true
	}
	return contain(root.Left, mem) || contain(root.Right, mem)
}
