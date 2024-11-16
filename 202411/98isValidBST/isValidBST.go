package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, mx, mi *TreeNode) bool
	dfs = func(node *TreeNode, mx, mi *TreeNode) bool {
		if node == nil {
			return true
		}
		// 这里是一个容易出错点，一定得是 >=
		if mx != nil && node.Val >= mx.Val {
			return false
		}
		// 一定得是 <=
		if mi != nil && node.Val <= mi.Val {
			return false
		}
		return dfs(node.Left, node, mi) && dfs(node.Right, mx, node)
	}
	return dfs(root, nil, nil)
}
