package main

import (
	"sort"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ans := help(root1)
	ans = append(ans, help(root2)...)
	sort.Ints(ans)
	return ans
}

func help(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	ans := help(node.Left)
	ans = append(ans, node.Val)
	ans = append(ans, help(node.Right)...)
	return ans
}
