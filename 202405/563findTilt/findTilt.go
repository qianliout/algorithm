package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func findTilt(root *TreeNode) int {
	sumMem := make(map[*TreeNode]int)
	tiltMem := make(map[*TreeNode]int)
	return tilt(root, tiltMem, sumMem)
}

func tilt(root *TreeNode, tiltMem, sumMem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := tiltMem[root]; ok {
		return va
	}
	left := sum(root.Left, sumMem)
	right := sum(root.Right, sumMem)

	ans := left - right
	if ans < 0 {
		ans = -ans
	}

	le := tilt(root.Left, tiltMem, sumMem)
	ri := tilt(root.Right, tiltMem, sumMem)
	ans = ans + le + ri

	tiltMem[root] = ans

	return ans
}

func sum(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}
	if root.Left == nil && root.Right == nil {
		mem[root] = root.Val
		return root.Val
	}
	left := sum(root.Left, mem)
	right := sum(root.Right, mem)
	mem[root] = root.Val + left + right

	return mem[root]
}
