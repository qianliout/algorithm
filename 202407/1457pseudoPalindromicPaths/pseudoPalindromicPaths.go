package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 给你一棵二叉树，每个节点的值为 1 到 9
func pseudoPalindromicPaths(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode)
	path := make([]int, 10)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path[node.Val]++

		if node.Left == nil && node.Right == nil {
			if check(path) {
				ans++
			}
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		path[node.Val]--
	}
	dfs(root)
	return ans
}

func check(path []int) bool {
	one := 0
	for _, ch := range path {
		if ch&1 == 1 {
			one++
		}
	}
	return one <= 1
}
