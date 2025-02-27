package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func findTargetNode2(root *TreeNode, cnt int) int {
	ans := dfs(root)
	n := len(ans)
	return ans[n-cnt]
}

func dfs(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	ans = append(ans, left...)
	ans = append(ans, root.Val)
	ans = append(ans, right...)
	return ans
}

// 反向中序遍历: 右 -> 根 -> 左
func findKthLargest(root *TreeNode, cnt int) int {
	c := 0
	res := -1
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		// 说了一定有答案
		if root == nil || c > cnt {
			return
		}
		helper(root.Right)
		c++
		if c == cnt {
			res = root.Val
			return
		}
		helper(root.Left)
	}
	helper(root)
	return res
}

// 反向中序遍历: 右 -> 根 -> 左
func findTargetNode(root *TreeNode, cnt int) int {
	var res int
	var helper func(node *TreeNode)
	c := 0 // 计数器

	helper = func(node *TreeNode) {
		if node == nil || c > cnt {
			return
		}
		// 先递归右子树
		helper(node.Right)
		// 访问当前节点
		c++
		if c == cnt {
			res = node.Val
			return
		}
		// 再递归左子树
		helper(node.Left)
	}
	helper(root)
	return res
}

func count(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}
	left := count(root.Left, mem)
	right := count(root.Right, mem)
	mem[root] = left + right + 1
	return left + right + 1
}
