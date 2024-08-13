package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func closestNodes(root *TreeNode, queries []int) [][]int {

	n := len(queries)
	nums := make([]int, 0)
	var dfs func(o *TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}
		dfs(o.Left)
		nums = append(nums, o.Val)
		dfs(o.Right)
	}
	dfs(root)
	ans := make([][]int, n)
	for i, ch := range queries {
		// mini 是树中小于等于 queries[i] 的 最大值 。如果不存在这样的值，则使用 -1 代替。
		// maxi 是树中大于等于 queries[i] 的 最小值 。如果不存在这样的值，则使用 -1 代替。
		a := findRight(nums, ch)
		b := findLeft(nums, ch)
		ans[i] = []int{a, b}
	}
	return ans
}

func findLeft(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if mid >= 0 && mid < len(nums) && nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l < 0 || l >= len(nums) || nums[l] < target {
		return -1
	}
	return nums[l]
}

func findRight(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l+1)>>1
		if mid >= 0 && mid < len(nums) && nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l >= len(nums) || nums[l] > target {
		return -1
	}
	return nums[l]
}

// 这样遍历会超时
func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inOrder(root.Left)
	left = append(left, root.Val)
	left = append(left, inOrder(root.Right)...)
	return left
}
