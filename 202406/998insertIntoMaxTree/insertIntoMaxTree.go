package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil || root.Val < val {
		node.Left = root
		return node
	}

	var pre *TreeNode
	cur := root
	for cur != nil && cur.Val > val {
		pre = cur
		cur = cur.Right
	}
	if pre == nil {
		node.Left = root
		return node
	}
	pre.Right = node
	node.Left = cur
	return root
}
