package main

import (
	"fmt"
	"sort"

	. "outback/algorithm/common/treenode"
)

func main() {
	fmt.Println(help([]int{7, 6, 8, 5}))
}

func minimumOperations(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		lev1 := make([]*TreeNode, 0)
		lev2 := make([]int, 0)
		for _, no := range queue {
			lev2 = append(lev2, no.Val)
			if no.Left != nil {
				lev1 = append(lev1, no.Left)
			}
			if no.Right != nil {
				lev1 = append(lev1, no.Right)
			}
		}
		ans += help(lev2)
		queue = lev1
	}

	return ans
}
func help(lev []int) int {
	n := len(lev)
	ans := 0
	id := make([]int, n) // 离散化后的数组
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return lev[id[i]] < lev[id[j]] })
	ans += n
	vis := make([]bool, n)
	for _, v := range id {
		if vis[v] {
			continue
		}
		for !vis[v] {
			vis[v] = true
			v = id[v]
		}
		ans--
	}
	return ans
}

// 最高级的做法
func help2(lev []int) int {
	n := len(lev)
	ans := 0
	id := make([]int, n) // 离散化后的数组
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return lev[id[i]] < lev[id[j]] })
	// 排序之后，id[i]里的值就表示 ans[i]的值应该在的位置
	// 交换次数是 n 减去 环的个数
	ans += n
	vis := make([]bool, n)

	for _, v := range id {
		if vis[v] {
			continue
		}
		// 说明有一个环了，那就把这个环里的所有都找完
		for ; !vis[v]; v = id[v] {
			vis[v] = true
		}
		ans--
	}
	return ans
}
