package main

import (
	"sort"
)

func main() {

}

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool { return happiness[i] > happiness[j] })
	sub := 0
	ans := 0
	for i := 0; i < k; i++ {
		ans += max(0, happiness[i]-sub)
		sub++
	}
	return int64(ans)
}
