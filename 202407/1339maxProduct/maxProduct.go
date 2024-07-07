package main

import (
	"math"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func maxProduct(root *TreeNode) int {
	var mod = int(math.Pow10(9)) + 7
	cnt := make(map[*TreeNode]int)
	allSum := sum(root, cnt)
	var dfs func(node *TreeNode)
	ans := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodeSum := sum(node, cnt)
		ans = max(ans, nodeSum*(allSum-nodeSum))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans % mod
}

func sum(node *TreeNode, cnt map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	if va, ok := cnt[node]; ok {
		return va
	}
	left := sum(node.Left, cnt)
	right := sum(node.Right, cnt)
	// 下面这几行可加可不加
	if node.Left != nil {
		cnt[node.Left] = left
	}
	if node.Right != nil {
		cnt[node.Right] = right
	}
	cnt[node] = left + right + node.Val
	return cnt[node]
}
