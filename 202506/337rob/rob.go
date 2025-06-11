package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 不加缓存会超时
func rob(root *TreeNode) int {
	mem := make(map[*TreeNode]int)
	ans := rob2(root, mem)
	return ans
}

func rob2(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}
	y := root.Val
	if root.Left != nil {
		y += rob2(root.Left.Left, mem) + rob2(root.Left.Right, mem)
	}
	if root.Right != nil {
		y += rob2(root.Right.Left, mem) + rob2(root.Right.Right, mem)
	}
	n := rob2(root.Left, mem) + rob2(root.Right, mem)
	mem[root] = max(y, n)
	return max(y, n)
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ly, ln := dfs(root.Left)
	ry, rn := dfs(root.Right)
	r := root.Val + ln + rn
	n := max(ly, ln) + max(ry, rn)
	return r, n
}
