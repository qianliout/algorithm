package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func treeQueries(root *TreeNode, queries []int) []int {
	height := make(map[*TreeNode]int)
	var dfs1 func(node *TreeNode) int // 算高度
	dfs1 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		res := 1 + max(dfs1(node.Left), dfs1(node.Right))
		height[node] = res
		return res
	}
	cnt := make(map[int]int)
	// 从删除node子树之后这棵树的深度
	var dfs2 func(node *TreeNode, dep, resH int)
	dfs2 = func(node *TreeNode, depth, resH int) {
		if node == nil {
			return
		}
		depth++
		cnt[node.Val] = resH
		dfs2(node.Left, depth, max(resH, depth+height[node.Right]))
		dfs2(node.Right, depth, max(resH, depth+height[node.Left]))
	}
	dfs1(root)
	dfs2(root, -1, 0)
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = cnt[queries[i]]
	}
	return ans
}
