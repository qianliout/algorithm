package main

import (
	"slices"
)

func main() {

}
func minimumTime(n int, relations [][]int, time []int) int {
	g := make([][]int, n) // g[i]表示完成i之后，还可以完成那些课程
	in := make([]int, n)  // in[i] 表示如果要完成 i，那么还有多少个先修课程要完成
	for _, ch := range relations {
		x, y := ch[0]-1, ch[1]-1
		g[x] = append(g[x], y)
		in[y]++
	}
	dp := make([]int, n)
	q := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		fir := q[0]
		q = q[1:]

		dp[fir] = dp[fir] + time[fir]

		for _, j := range g[fir] {
			dp[j] = max(dp[j], dp[fir])
			in[j]--
			if in[j] == 0 {
				q = append(q, j)
			}
		}
	}
	return slices.Max(dp)
}
