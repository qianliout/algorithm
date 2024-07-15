package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
}

func bestTeamScore(scores []int, ages []int) int {
	pairs := make([]pair, 0)
	n := len(scores)
	for i := 0; i < n; i++ {
		pairs = append(pairs, pair{
			sco: scores[i],
			age: ages[i],
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].sco < pairs[j].sco {
			return true
		} else if pairs[i].sco > pairs[j].sco {
			return false
		}
		return pairs[i].age <= pairs[j].age
	})
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = pairs[i].sco
		for j := i - 1; j >= 0; j-- {
			if pairs[j].age <= pairs[i].age {
				dp[i] = max(dp[i], dp[j]+pairs[i].sco)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

type pair struct {
	sco int
	age int
}
