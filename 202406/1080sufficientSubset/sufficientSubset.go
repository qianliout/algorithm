package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit -= root.Val
	if root.Left == nil && root.Right == nil {
		if limit > 0 {
			return nil
		}
		return root
	}
	if root.Left != nil {
		root.Left = sufficientSubset(root.Left, limit)
	}

	if root.Right != nil {
		root.Right = sufficientSubset(root.Right, limit)
	}
	// 这里是最不好理解的
	// 如果 node 的儿子都被删除，说明经过 node 的所有儿子的路径和都小于 limit，这等价于经过 node 的所有路径和都小于 limit，所以 node 需要被删除。
	if root.Left != nil || root.Right != nil {
		return root
	}
	return nil
}
