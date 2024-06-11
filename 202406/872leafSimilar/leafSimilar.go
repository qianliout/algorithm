package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	res1 := dfs(root1)
	res2 := dfs(root2)
	if len(res1) != len(res2) {
		return false
	}
	for i := 0; i < len(res1); i++ {
		if res1[i] != res2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	if root.Left == nil && root.Right == nil {
		res = append(res, root.Val)
		return res
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res = append(res, left...)
	res = append(res, right...)
	return res
}
