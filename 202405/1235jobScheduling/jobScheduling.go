package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	type pair struct{ start, end, profit int }
	work := make([]pair, n)
	for i := range startTime {
		work[i] = pair{start: startTime[i], end: endTime[i], profit: profit[i]}
	}
	sort.Slice(work, func(i, j int) bool { return work[i].end <= work[j].end })

	dp := make([]int, n+1)
	for i, w := range work {
		dp[i+1] = max(dp[i+1], dp[i]) // 不做
		p := sort.Search(i, func(j int) bool { return work[j].end > w.start }) - 1
		dp[i+1] = max(dp[i+1], dp[p+1]+w.profit) // 做
	}
	return dp[n]
}
