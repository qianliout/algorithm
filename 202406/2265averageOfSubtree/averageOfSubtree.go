package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func averageOfSubtree(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	ans := 0
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		ls, lc := dfs(node.Left)
		rs, rc := dfs(node.Right)

		sum := ls + rs + node.Val
		cnt := lc + rc + 1
		if sum/cnt == node.Val {
			ans++
		}
		return sum, cnt
	}

	dfs(root)
	return ans
}
