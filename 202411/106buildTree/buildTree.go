package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) != len(postorder) {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	idx := findIndex(inorder, postorder[len(postorder)-1])
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
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
