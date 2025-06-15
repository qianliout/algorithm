package main

import (
	. "outback/algorithm/common/listnode"
	. "outback/algorithm/common/treenode"
)

func main() {

}

func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 0
	}
	st := make([]int, 0)
	for i, c := range nums {
		for len(st) > 0 && c > nums[st[len(st)-1]] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			ans[last] = c
		}
		st = append(st, i)
	}
	return ans
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mx := findMax(nums)
	root := &TreeNode{Val: nums[mx]}
	root.Left = constructMaximumBinaryTree(nums[:mx])
	root.Right = constructMaximumBinaryTree(nums[mx+1:])
	return root
}

func findMax(nums []int) int {
	mx := 0
	for i, c := range nums {
		if c > nums[mx] {
			mx = i
		}
	}
	return mx
}
