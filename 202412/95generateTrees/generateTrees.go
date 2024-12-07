package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func generateTrees(n int) []*TreeNode {

	var dfs func(i, j int) []*TreeNode
	dfs = func(i, j int) []*TreeNode {
		if i > j {
			return []*TreeNode{nil}
		}
		if i == j {
			return []*TreeNode{&TreeNode{Val: j}}
		}
		ans := make([]*TreeNode, 0)
		for k := i; k <= j; k++ {
			left := dfs(i, k-1)
			right := dfs(k+1, j)
			for _, le := range left {
				for _, ri := range right {
					ans = append(ans, &TreeNode{Left: le, Right: ri, Val: k})
				}
			}
		}
		return ans
	}
	ans := dfs(1, n)
	return ans
}
