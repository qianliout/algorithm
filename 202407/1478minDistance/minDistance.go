package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDistance([]int{1, 4, 8, 10, 20}, 3))
}

func minDistance(houses []int, k int) int {
	n := len(houses)
	sort.Ints(houses)
	inf := math.MaxInt / 10
	ms := make([][]int, n)
	for i := range ms {
		ms[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			mid := (i + j) / 2
			for m := i; m <= j; m++ {
				ms[i][j] += abs(houses[mid] - houses[m])
			}
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	for i := 0; i < n; i++ {
		dp[i][1] = ms[0][i]
		for m := 2; m <= k; m++ {
			for j := 0; j < i; j++ {
				dp[i][m] = min(dp[i][m], dp[j][m-1]+ms[j+1][i])
			}
		}
	}

	return dp[n-1][k]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
