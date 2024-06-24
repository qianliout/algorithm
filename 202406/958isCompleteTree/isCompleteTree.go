package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

// 完全二叉树的一个特性，按层序遍历，第一个空节点之后再没有其他节点
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}

	findNil := false
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no == nil {
				findNil = true
			} else {
				if findNil {
					return false
				}
				lev = append(lev, no.Left)
				lev = append(lev, no.Right)
			}
		}
		queue = lev
	}
	return true
}
