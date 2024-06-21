package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	faMap := make(map[*TreeNode]*TreeNode)
	// 每个节点值都不相同，所以才这样做
	visit := make(map[int]bool)
	var dfs1 func(root *TreeNode)
	dfs1 = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			faMap[root.Left] = root
			dfs1(root.Left)
		}
		if root.Right != nil {
			faMap[root.Right] = root
			dfs1(root.Right)
		}
	}
	dfs1(root)

	// 再一次 dfs 找节点
	path := make([]int, 0)
	var dfs2 func(root *TreeNode, k int)

	dfs2 = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		if visit[root.Val] {
			return
		}

		visit[root.Val] = true
		if k == 0 {
			path = append(path, root.Val)
			return
		}
		dfs2(root.Left, k-1)
		dfs2(root.Right, k-1)
		dfs2(faMap[root], k-1)
	}
	dfs2(target, k)

	return path
}

// 参考的题解：https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/solutions/900233/863-er-cha-shu-zhong-suo-you-ju-chi-wei-far5x/

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
