package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 1))
}

func maxJumps(arr []int, d int) int {
	n := len(arr)
	pairs := make([]pair, n)
	for i, ch := range arr {
		pairs[i] = pair{Idx: i, H: ch}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].H < pairs[j].H {
			return true
		} else if pairs[i].H > pairs[j].H {
			return false
		} else {
			return pairs[i].Idx <= pairs[j].Idx
		}
	})
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for _, ch := range pairs {
		idx, h := ch.Idx, ch.H
		// 向左跳
		cur := 1
		for i := idx - 1; i >= 0 && i >= idx-d; i-- {
			if arr[i] >= h {
				break
			}
			cur = max(dp[i]+1, cur)
		}
		// 向右跳
		for i := idx + 1; i < n && i <= idx+d; i++ {
			if arr[i] >= h {
				break
			}
			cur = max(dp[i]+1, cur)
		}
		dp[idx] = max(dp[idx], cur)
	}

	for i := 0; i < n; i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

type pair struct {
	H   int
	Idx int
}
