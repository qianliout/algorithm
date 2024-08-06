package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {
	createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}})
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	// 值不同
	nodeM := make(map[int]*TreeNode)
	childM := make(map[int]bool)

	for _, ch := range descriptions {
		p, c, l := ch[0], ch[1], ch[2]
		if nodeM[p] == nil {
			nodeM[p] = &TreeNode{Val: p}
		}
		if nodeM[c] == nil {
			nodeM[c] = &TreeNode{Val: c}
		}
		childM[c] = true
		// 如果 isLefti == 1 ，那么 childi 就是 parenti 的左子节点。
		// 如果 isLefti == 0 ，那么 childi 就是 parenti 的右子节点。
		if l == 1 {
			nodeM[p].Left = nodeM[c]
		} else {
			nodeM[p].Right = nodeM[c]
		}
	}
	for k, v := range nodeM {
		if !childM[k] {
			return v
		}
	}
	return nil
}
