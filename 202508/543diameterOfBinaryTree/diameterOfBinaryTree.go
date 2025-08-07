package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	ans := 0
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)  // 左子树的最大深度（节点数）
		r := dfs(node.Right) // 右子树的最大深度（节点数）

		// 经过当前节点的路径长度（节点数）= 左深度 + 右深度 + 1
		ans = max(ans, l+r+1)

		// 返回以当前节点为根的子树的最大深度（节点数）
		return max(l, r) + 1
	}
	dfs(root)
	// 转换为边数：节点数 - 1 = 边数
	return ans - 1
}

/*
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
*/
