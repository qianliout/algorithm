package main

import (
	"sort"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func minimumOperations(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		lev1 := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Left != nil {
				lev1 = append(lev1, no.Left)
			}
			if no.Right != nil {
				lev1 = append(lev1, no.Right)
			}
		}
		a, b := help(lev1)
		ans += b
		queue = a
	}

	return ans
}

// 这个做法是交换结点，题目中说的是交换节点的值，所以不对
func help(lev []*TreeNode) ([]*TreeNode, int) {
	exit := make(map[int]*TreeNode)
	ans := make([]int, len(lev))
	for i := range lev {
		exit[lev[i].Val] = lev[i]
		ans[i] = lev[i].Val
	}
	cnt := 0
	res := make([]*TreeNode, 0)
	sort.Ints(ans)
	for i := range ans {
		if ans[i] != lev[i].Val {
			cnt++
		}
		res = append(res, exit[ans[i]])
	}

	return res, cnt / 2
}
