package main

import (
	. "outback/algorithm/common/listnode"
	. "outback/algorithm/common/treenode"
)

func main() {

}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if check(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func check(list *ListNode, tree *TreeNode) bool {
	if list == nil {
		return true
	}
	if tree == nil {
		return false
	}
	if list.Val != tree.Val {
		return false
	}

	return check(list.Next, tree.Left) || check(list.Next, tree.Right)
}
