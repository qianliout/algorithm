package main

import (
	"sort"
)

func main() {

}

func maximumBags(capacity []int, rocks []int, k int) int {
	n := len(capacity)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = capacity[i] - rocks[i]
	}
	sort.Ints(diff)
	ans := 0
	for i := 0; i < n && k > 0; i++ {
		if k >= diff[i] {
			ans++
			k -= diff[i]
		}
	}
	return ans
}
