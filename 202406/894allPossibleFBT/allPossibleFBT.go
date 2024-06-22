package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func allPossibleFBT(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{}}
	}
	if n%2 == 0 {
		return []*TreeNode{}
	}

	ans := make([]*TreeNode, 0)
	for i := 1; i < n; i++ {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)
		for _, l := range left {
			for _, r := range right {
				node := &TreeNode{Left: l, Right: r, Val: 0}
				ans = append(ans, node)
			}
		}
	}
	return ans
}
