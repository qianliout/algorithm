package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestPathValue("abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}))
}

func largestPathValue(colors string, edges [][]int) int {

	n := len(colors)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	g := make([][]int, n)

	in := make([]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		in[y]++
	}
	queue := make([]int, 0)
	for k, v := range in {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	cnt := 0
	visit := make([]bool, n)
	for len(queue) > 0 {
		cnt++
		x := queue[0]
		queue = queue[1:]
		if visit[x] {
			return -1
		}
		visit[x] = true

		col := int(colors[x]) - int('a')
		dp[x][col]++

		for _, nex := range g[x] {
			for i := 0; i < 26; i++ {
				dp[nex][i] = max(dp[nex][i], dp[x][i])
			}
			in[nex]--
			if in[nex] == 0 {
				queue = append(queue, nex)
			}
		}
	}
	// 判断是否有环
	if cnt < n {
		return -1
	}
	ans := 0

	for i := range dp {
		for _, ch := range dp[i] {
			ans = max(ans, ch)
		}
	}
	return ans
}
