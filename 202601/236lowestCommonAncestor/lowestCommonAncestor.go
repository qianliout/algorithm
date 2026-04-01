package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right != nil {
		return right
	}
	if left != nil && right == nil {
		return left
	}
	if left != nil && right != nil {
		return root
	}
	return nil
}
