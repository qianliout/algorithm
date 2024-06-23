package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	a := flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) // 都不翻转或都翻转
	b := flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
	return a || b
}
