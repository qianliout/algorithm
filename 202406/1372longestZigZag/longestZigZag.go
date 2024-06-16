package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 可以得到解
func longestZigZag1(root *TreeNode) int {
	var ans int
	check(root, &ans)
	return ans
}

func check(root *TreeNode, ans *int) (int, int) {
	if root == nil {
		return 0, 0
	}
	_, leftRight := check(root.Left, ans)
	rightLeft, _ := check(root.Right, ans)

	*ans = max(*ans, max(leftRight, rightLeft))
	return leftRight + 1, rightLeft + 1
}

func longestZigZag(root *TreeNode) int {
	var ans int
	dfs(root, 0, 0, &ans)
	return ans
}

// 这重解法好理解一点
// l 表示从 l这边来一共走了多少边
func dfs(root *TreeNode, l, r int, ans *int) int {

	*ans = max(*ans, l, r)
	if root == nil {
		return 0
	}
	if root.Left != nil {
		dfs(root.Left, r+1, 0, ans)
	}
	if root.Right != nil {
		dfs(root.Right, 0, l+1, ans)
	}
	return 1
}
