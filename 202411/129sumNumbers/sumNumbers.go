package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func sumNumbers1(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			sum += num
			return
		}
		num = num*10 + node.Val
		dfs(node.Left, num)
		dfs(node.Right, num)
	}
	dfs(root, 0)
	// 到root==nil 才计算结果，会计算两次，所以除以2，但是这样是不对的,因为如果一个节点只有一个nil 那就只能计算一次，
	return sum / 2
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, num int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		num = num*10 + node.Val
		if node.Right == nil && node.Left == nil {
			sum += num
			return
		}
		dfs(node.Left, num)
		dfs(node.Right, num)
	}
	dfs(root, 0)
	return sum
}
