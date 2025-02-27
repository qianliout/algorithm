package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	mem := make(map[*TreeNode]int)
	left := dep(root.Left, mem)
	right := dep(root.Right, mem)
	if abs(left-right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dep(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}

	left := dep(root.Left, mem)
	right := dep(root.Right, mem)

	mem[root] = max(left, right) + 1
	return mem[root]
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	return left
}
