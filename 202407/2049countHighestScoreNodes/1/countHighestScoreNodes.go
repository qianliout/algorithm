package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
	// fmt.Println(countHighestScoreNodes([]int{-1, 0}))
	// fmt.Println(countHighestScoreNodes([]int{-1, 3, 3, 5, 7, 6, 0, 0}))
	// fmt.Println(countHighestScoreNodes([]int{-1, 12, 4, 0, 8, 7, 4, 8, 0, 14, 12, 14, 3, 7, 3}))
}

func countHighestScoreNodes(parents []int) int {
	root := &TreeNode{Val: 0}
	cnt := make(map[int]*TreeNode)
	cnt[0] = root
	n := len(parents)
	// parents[i] 是节点 i 的父节点
	for i := 1; i < n; i++ {
		node := &TreeNode{Val: i}
		cnt[i] = node
	}
	// 构造二叉树
	for i, p := range parents {
		if p == -1 {
			continue
		}
		fa := cnt[p]
		chi := cnt[i]
		if fa.Left == nil {
			fa.Left = chi
		} else if fa.Right == nil {
			fa.Right = chi
		}
	}

	// 找node 的所了子节点之和
	mem := make(map[int]int)
	// 找一个节点的所有子节点的节点数(包括自已)
	var dfs1 func(node *TreeNode) int
	dfs1 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if va, ok := mem[node.Val]; ok {
			return va
		}
		le := dfs1(node.Left)
		ri := dfs1(node.Right)
		mem[node.Val] = le + ri + 1
		return mem[node.Val]
	}
	dfs1(root)

	ans := 0
	mx := 0

	// 再一次 dfs 找答案
	var dfs2 func(node *TreeNode)
	dfs2 = func(node *TreeNode) {
		if node == nil {
			return
		}
		le := dfs1(node.Left)
		ri := dfs1(node.Right)
		remain := n - le - ri - 1
		score := max(1, le) * max(1, remain) * max(1, ri)
		if score == mx {
			ans++
		} else if score > mx {
			mx = score
			ans = 1
		}
		dfs2(node.Left)
		dfs2(node.Right)

	}
	dfs2(root)

	return ans
}
