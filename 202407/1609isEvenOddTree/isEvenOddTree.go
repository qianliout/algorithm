package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil || root.Val&1 == 1 {
		return false
	}
	queue := []*TreeNode{root}
	dep := 0
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		if len(lev) > 0 {
			dep++
			if dep&1 == 1 && !check1(lev) {
				return false
			}
			if dep&1 == 0 && !check2(lev) {
				return false
			}
			queue = lev
		}
	}

	return true
}

// 奇数层
// 奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
func check1(lev []*TreeNode) bool {
	for i := 0; i < len(lev); i++ {
		if lev[i].Val&1 == 1 {
			return false
		}
		if i > 0 && lev[i].Val >= lev[i-1].Val {
			return false
		}

	}
	return true
}

// 偶数层
// 偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
func check2(lev []*TreeNode) bool {
	for i := 0; i < len(lev); i++ {
		if lev[i].Val&1 == 0 {
			return false
		}
		if i > 0 && lev[i].Val <= lev[i-1].Val {
			return false
		}

	}
	return true
}
