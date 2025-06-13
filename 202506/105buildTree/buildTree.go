package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	idx := find(inorder, preorder[0])
	root.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree(preorder[1+idx:], inorder[idx+1:])
	return root
}

func find(order []int, va int) int {
	for i, c := range order {
		if c == va {
			return i
		}
	}
	return 0
}
