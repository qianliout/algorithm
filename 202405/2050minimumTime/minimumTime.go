package main

import (
	"fmt"

	. "outback/algorithm/common/utils"
)

func main() {
	fmt.Println(minimumTime(3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}))
}

func minimumTime(n int, relations [][]int, time []int) int {

	g := make([][]int, n)
	in := make([]int, n)
	for _, ch := range relations {
		x, y := ch[0]-1, ch[1]-1
		g[x] = append(g[x], y)
		in[y]++
	}

	queue := make([]int, 0)
	for k, v := range in {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	dp := make([]int, n)
	for len(queue) > 0 {
		// 出队，意味着所有先修课都上完了
		fir := queue[0]
		queue = queue[1:]

		dp[fir] = dp[fir] + time[fir] // 说明之前的课程都学习完了，加上本次学习的课程,就是学习到 fir 这个课程的所花时间

		for _, nex := range g[fir] {
			// fir 是 nex 的先修课程,所以这一步就是计算 nex 的所有先修课程中的最大值
			dp[nex] = max(dp[fir], dp[nex])

			in[nex]--
			if in[nex] == 0 {
				queue = append(queue, nex)
			}
		}
	}

	return Max(dp...)
}
