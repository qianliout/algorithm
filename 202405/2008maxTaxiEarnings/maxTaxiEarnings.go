package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTaxiEarnings(5, [][]int{{2, 5, 4}, {1, 5, 1}}))
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool { return rides[i][1] <= rides[j][1] })

	dp := make([]int64, n+1)
	for i, ri := range rides {
		dp[i+1] = max(dp[i+1], dp[i])
		p := sort.Search(i, func(j int) bool { return rides[j][1] > ri[0] }) - 1
		dp[i+1] = max(dp[i+1], dp[p+1]+int64(ri[1]-ri[0]+ri[2]))
	}
	return dp[len(rides)]
}
