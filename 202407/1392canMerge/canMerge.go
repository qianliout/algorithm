package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func canMerge(trees []*TreeNode) *TreeNode {
	leaf := make(map[int]*TreeNode)
	all := make(map[int]*TreeNode)
	// 所有的叶子值节点不重复
	for _, no := range trees {
		if no == nil {
			continue
		}
		all[no.Val] = no

		if no.Left != nil {
			if leaf[no.Left.Val] != nil {
				return nil
			}
			leaf[no.Left.Val] = no.Left
		}
		if no.Right != nil {
			if leaf[no.Right.Val] != nil {
				return nil
			}
			leaf[no.Right.Val] = no.Right
		}
	}

	// 叶子节点集合是leaf,那么只有一个根节点不在 leaf 中这个节点就是合并后的树的根节点
	var root *TreeNode

	exist := make(map[int]int)

	for _, no := range trees {
		if leaf[no.Val] == nil {
			if root != nil {
				return nil
			}
			root = no
		}
		if exist[no.Val] > 0 {
			return nil
		}
		exist[no.Val]++
	}
	if root == nil {
		return nil
	}
	// 验证完成之后，把根节点也加入的节点集合中，用于合并时找到对应的节点
	for k, v := range all {
		leaf[k] = v
	}
	// 开始合并
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root.Left != nil {
			if leaf[root.Left.Val] != nil {
				root.Left = leaf[root.Left.Val]
				leaf[root.Left.Val] = nil
				dfs(root.Left)
			}
		}
		if root.Right != nil {
			if leaf[root.Right.Val] != nil {
				root.Right = leaf[root.Right.Val]
				leaf[root.Right.Val] = nil
				dfs(root.Right)
			}
		}
	}
	// 做一步的目的是下面验证所有节点是否都已用完
	leaf[root.Val] = nil

	dfs(root)

	// 看所有的节点是否都用完
	for _, v := range leaf {
		if v != nil {
			return nil
		}
	}
	// 最后验证一下是不是真正的二叉搜索树
	if checkBinaryTree(root) {
		return root
	}

	return nil
}

func checkBinaryTree(root *TreeNode) bool {
	var dfs func(root, mx, mi *TreeNode) bool
	dfs = func(root, mx, mi *TreeNode) bool {
		if root == nil {
			return true
		}
		if mx != nil && root.Val >= mx.Val {
			return false
		}
		if mi != nil && root.Val <= mi.Val {
			return false
		}
		return dfs(root.Left, root, mi) && dfs(root.Right, mx, root)
	}
	return dfs(root, nil, nil)
}
