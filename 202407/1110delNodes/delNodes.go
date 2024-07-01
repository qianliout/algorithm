package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func delNodes(root *TreeNode, to []int) []*TreeNode {
	del := make(map[int]bool)
	for _, x := range to {
		del[x] = true
	}
	ans := make([]*TreeNode, 0)

	// dfs表示对 node 执行删除操作的节点值
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if !del[node.Val] {
			return node
		}
		if node.Left != nil {
			ans = append(ans, node.Left)
		}
		if node.Right != nil {
			ans = append(ans, node.Right)
		}

		return nil
	}
	if dfs(root) != nil {
		ans = append(ans, root)
	}
	return ans
}
