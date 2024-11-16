package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	flatten(left)
	flatten(right)

	root.Right = left
	root.Left = nil
	cur := root
	for cur != nil && cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = right
	cur.Left = nil
}

// 展开到右子树上去
