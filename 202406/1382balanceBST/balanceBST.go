package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	pre(root, &nums)
	return gen(nums)
}

func gen(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	n := len(nums)
	mid := n / 2
	left := nums[:mid]
	right := nums[mid+1:]

	return &TreeNode{
		Val:   nums[mid],
		Left:  gen(left),
		Right: gen(right),
	}
}

func pre(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	pre(root.Left, res)
	*res = append(*res, root.Val)
	pre(root.Right, res)
}
