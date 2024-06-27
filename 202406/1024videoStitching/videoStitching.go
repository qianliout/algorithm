package main

import (
	"sort"
)

func main() {

}

func videoStitching(clips [][]int, t int) int {
	base := 101 // 最多100个区间
	sort.Slice(clips, func(i, j int) bool { return clips[i][0] < clips[j][0] })

	// dp[i] 表示[0,i]的结果，注意闭区间
	dp := make([]int, t+1)
	for i := range dp {
		dp[i] = base
	}
	dp[0] = 0
	for _, ch := range clips {
		start, end := ch[0], ch[1]
		// 只要[1:t] 这个区间，但是 clips里可能超过这个区间
		for j := start + 1; j <= min(end, t); j++ {
			dp[j] = min(dp[j], dp[start]+1)
		}
	}
	if dp[t] >= base {
		return -1
	}
	return dp[t]
}
