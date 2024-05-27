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
		fir := queue[0]
		queue = queue[1:]

		dp[fir] = dp[fir] + time[fir] // 说明之前的课程都学习完了，加上本次学习的课程就是这一条线的全部时间

		for _, nex := range g[fir] {
	
			dp[nex] = max(dp[fir], dp[nex])

			in[nex]--
			if in[nex] == 0 {
				queue = append(queue, nex)
			}
		}
	}

	return Max(dp...)
}
