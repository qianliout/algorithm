package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func deduceTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	idx := findIdx(inorder, preorder[0])
	node.Left = deduceTree(preorder[1:idx+1], inorder[:idx])
	node.Right = deduceTree(preorder[idx+1:], inorder[idx+1:])
	return node
}

// 注意：preorder 和 inorder 中均不含重复数字。

func findIdx(nums []int, va int) int {
	for i, ch := range nums {
		if ch == va {
			return i
		}
	}
	return len(nums)
}
