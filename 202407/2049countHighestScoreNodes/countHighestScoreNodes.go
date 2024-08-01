package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
	fmt.Println(countHighestScoreNodes([]int{-1, 0}))
	fmt.Println(countHighestScoreNodes([]int{-1, 3, 3, 5, 7, 6, 0, 0}))
	fmt.Println(countHighestScoreNodes([]int{-1, 12, 4, 0, 8, 7, 4, 8, 0, 14, 12, 14, 3, 7, 3}))
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
		pn := cnt[p]
		in := cnt[i]
		if pn.Left == nil {
			pn.Left = in
		} else if pn.Right == nil {
			pn.Right = in
		}
	}
	// 找node 的所了子节点之和
	mem := make(map[int]int)
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if va, ok := mem[node.Val]; ok {
			return va
		}
		le := dfs(node.Left)
		ri := dfs(node.Right)
		mem[node.Val] = le + ri + 1
		return mem[node.Val]
	}
	dfs(root)
	ans := 0
	mx := 0
	for i := 0; i < n; i++ {
		node := cnt[i]
		le := dfs(node.Left)
		ri := dfs(node.Right)
		pr := 1
		if le == 0 && ri == 0 {
			le = 1
			ri = max(1, n-1)
		} else if le == 0 && ri != 0 {
			le = max(1, n-1-ri)
		} else if ri == 0 && le != 0 {
			ri = max(1, n-1-le)
		} else {
			pr = n - le - ri
		}

		a := le * ri * pr

		if a == mx {
			ans++
		} else if a > mx {
			mx = a
			ans = 1
		}
	}
	return ans
}
