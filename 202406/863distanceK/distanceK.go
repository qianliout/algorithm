package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

}

func isMonotonic(nums []int) bool {
	return add(nums) || sub(nums)
}

func add(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func sub(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}

func increasingBST(root *TreeNode) *TreeNode {
	ans := inorder(root)
	Dump := &TreeNode{}
	cur := Dump
	for i := 0; i < len(ans); i++ {
		cur.Right = &TreeNode{Val: ans[i]}
		cur = cur.Right
	}
	return Dump.Right
}

func inorder(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	left := inorder(root.Left)
	right := inorder(root.Right)
	ans = append(ans, left...)
	ans = append(ans, root.Val)
	ans = append(ans, right...)
	return ans
}
