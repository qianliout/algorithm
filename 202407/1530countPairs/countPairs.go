package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func countPairs(root *TreeNode, distance int) int {
	var ret int

	var dfs func(*TreeNode) []int

	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		if node.Left == nil && node.Right == nil {
			return []int{0} // Leaf node
		}

		left := append([]int(nil), addOne(dfs(node.Left))...)
		right := append([]int(nil), addOne(dfs(node.Right))...)

		for _, l := range left {
			for _, r := range right {
				if l+r <= distance {
					ret++
				}
			}
		}

		return append(left, right...)
	}

	dfs(root)
	return ret
}

func addOne(distances []int) []int {
	for i := range distances {
		distances[i]++
	}
	return distances
}
