package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(largestPathValue("abcd", [][]int{{0, 1}, {1, 2}, {3, 0}})) // 3
}

// 错误的解法
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
	}

	var dfs func(x int, col []int, vist []int) int

	dfs = func(x int, col []int, vist []int) int {
		if vist[x] > 0 {
			return -1
		}
		vist[x]++
		c := int(colors[x] - 'a')
		col[c]++
		for _, y := range g[x] {
			ans := dfs(y, col, vist)
			if ans < 0 {
				return -1
			}
		}
		return slices.Max(col)
	}

	ans := 0
	for i := 0; i < n; i++ {
		col := make([]int, 26)
		vist := make([]int, n)
		r := dfs(i, col, vist)
		if r == -1 {
			return -1
		}
		ans = max(ans, r)
	}
	return ans
}

/*
给你一个 有向图 ，它含有 n 个节点和 m 条边。节点编号从 0 到 n - 1 。
给你一个字符串 colors ，其中 colors[i] 是小写英文字母，表示图中第 i 个节点的 颜色 （下标从 0 开始）。同时给你一个二维数组 edges ，其中 edges[j] = [aj, bj] 表示从节点 aj 到节点 bj 有一条 有向边 。
图中一条有效 路径 是一个点序列 x1 -> x2 -> x3 -> ... -> xk ，对于所有 1 <= i < k ，从 xi 到 xi+1 在图中有一条有向边。路径的 颜色值 是路径中 出现次数最多 颜色的节点数目。
请你返回给定图中有效路径里面的 最大颜色值 。如果图中含有环，请返回 -1 。
*/
