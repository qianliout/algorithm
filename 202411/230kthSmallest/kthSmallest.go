package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	nums := make([]int, 0)
	var dfs1 func(node *TreeNode)
	dfs1 = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs1(node.Left)
		nums = append(nums, node.Val)
		dfs1(node.Right)
	}
	dfs1(root)
	if len(nums) < k {
		return 0
	}
	return nums[k-1]
}
