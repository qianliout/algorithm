package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {
	root := &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
	bstToGst(root)
	PreOrderTraversal(root)
}

func bstToGst2(root *TreeNode) *TreeNode {
	nums, nodes := pre(root)

	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		nums[i] += nums[i+1]
	}
	// 重新赋值
	for i := 0; i < len(nodes); i++ {
		nodes[i].Val = nums[i]
	}
	return root
}

func pre(node *TreeNode) ([]int, []*TreeNode) {
	if node == nil {
		return []int{}, []*TreeNode{}
	}
	left1, left2 := pre(node.Left)
	right1, right2 := pre(node.Right)
	nodes := make([]*TreeNode, 0)
	res := make([]int, 0)
	res = append(res, left1...)
	res = append(res, node.Val)
	res = append(res, right1...)

	nodes = append(nodes, left2...)
	nodes = append(nodes, node)
	nodes = append(nodes, right2...)

	return res, nodes
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}
