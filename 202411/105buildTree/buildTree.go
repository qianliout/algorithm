package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	idx := findIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}

func findIndex(nums []int, n int) int {
	for i, ch := range nums {
		if ch == n {
			return i
		}
	}
	return -1
}

// 无重复元素
