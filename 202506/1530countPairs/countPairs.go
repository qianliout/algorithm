package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

/*
给你二叉树的根节点 root 和一个整数 distance 。
如果二叉树中两个 叶 节点之间的 最短路径长度 小于或者等于 distance ，那它们就可以构成一组 好叶子节点对 。
返回树中 好叶子节点对的数量 。
*/
func countPairs(root *TreeNode, distance int) int {
	var count int
	var dfs func(node *TreeNode) []int
	// DFS 函数返回当前节点到其所有叶子节点的距离列表，然后在每个内部节点处统计左右子树叶子节点之间的有效配对。
	dfs = func(node *TreeNode) []int {
		if node == nil {
			// 空节点：返回空数组
			return []int{}
		}
		if node.Left == nil && node.Right == nil {
			// 当前节点是叶子节点了，自已到自己的距离是0
			return []int{0}
		}
		// 递归获取左子树中所有叶子节点到左子树根节点的距离
		// 递归获取右子树中所有叶子节点到右子树根节点的距离
		// 使用 addOne 函数将距离都加 1（因为要经过当前节点）
		left := dfs(node.Left)
		right := dfs(node.Right)
		left = addOne(left)
		right = addOne(right)
		// 遍历左子树的所有叶子节点距离 l
		// 遍历右子树的所有叶子节点距离 r
		// 如果 l + r <= distance，说明这对叶子节点满足条件，计数器加 1
		for _, l := range left {
			for _, r := range right {
				if l+r <= distance {
					count++
				}
			}
		}

		left = append(left, right...)
		return left
	}
	dfs(root)
	return count
}

func addOne(distances []int) []int {
	for i := range distances {
		distances[i] = distances[i] + 1
	}
	return distances
}
