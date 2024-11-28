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
		return envs[i].y < envs[j].y
	})
	f := make([]int, n)
	ans := 0
	for i, ch := range envs {
		f[i] = 1
		for j := i - 1; j >= 0; j-- {
			if ch.y > envs[j].y && ch.x > envs[j].x {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}

type envelope struct {
	x, y int
}
