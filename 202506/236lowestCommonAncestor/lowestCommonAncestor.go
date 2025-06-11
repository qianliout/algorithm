package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 把root理解成当前正在审查的节点
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		//  说明p和q分别在左右子树中找到
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
