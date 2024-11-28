package main

import (
	"sort"
)

func main() {

}

// 不用二分会超时
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	envs := make([]envelope, n)
	for i := 0; i < n; i++ {
		x, y := envelopes[i][0], envelopes[i][1]
		envs[i] = envelope{x, y}
	}
	sort.Slice(envs, func(i, j int) bool {
		if envs[i].x != envs[j].x {
			return envs[i].x < envs[j].x
		}
		return envs[i].y > envs[j].y
	})
	dp := make([]int, 0)

	for _, ch := range envs {
		idx := sort.SearchInts(dp, ch.y) - 1
		if idx == len(dp) {
			dp = append(dp, idx)
		} else {
			dp[idx] = ch.y
		}
	}
	return len(dp)
}

type envelope struct {
	x, y int
}
