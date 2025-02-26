package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 这样写的问题是会把结果计算两次
func pathSum2(root *TreeNode, targetSum int) [][]int {
	var dfs func(node *TreeNode, path []int, sum int)
	ans := make([][]int, 0)

	dfs = func(node *TreeNode, path []int, sum int) {
		if sum == targetSum && len(path) > 0 && node == nil {
			ans = append(ans, append([]int{}, path...))
			// 不能返，因为可能有负数
		}
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum += node.Val
		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
		path = path[:len(path)-1]
		sum -= node.Val
	}
	dfs(root, []int{}, 0)
	return ans
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(node *TreeNode, path []int, sum int)
	ans := make([][]int, 0)

	dfs = func(node *TreeNode, path []int, sum int) {

		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			ans = append(ans, append([]int{}, path...))
		}

		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
		path = path[:len(path)-1]
		sum -= node.Val
	}
	dfs(root, []int{}, 0)
	return ans
}

// 从根节点到叶子节点 路径总和等于给定目标和的路径
