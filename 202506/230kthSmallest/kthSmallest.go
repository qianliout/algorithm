package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	left := count(root.Left)

	if left+1 == k {
		return root.Val
	}
	if left >= k {
		return kthSmallest(root.Left, k)
	}
	return kthSmallest(root.Right, k-left-1)
}

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := count(root.Left)
	right := count(root.Right)
	return left + right + 1
}
