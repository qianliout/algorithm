package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 中序遍历
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var dfs func(root *TreeNode) bool

	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		// 先遍历左子树
		if !dfs(root.Left) {
			return false
		}
		// 比较当前 root
		if pre != nil && pre.Val >= root.Val {
			return false
		}
		// 对于右子树，前一个节点就是 root
		pre = root
		// 遍历右子树
		ans := dfs(root.Right)
		return ans
	}

	ans := dfs(root)
	return ans
}
