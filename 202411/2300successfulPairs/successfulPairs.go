package main

import (
	"sort"
)

func main() {

}

func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(spells)
	m := len(potions)
	sort.Ints(potions)
	ans := make([]int, n)
	for i, ch := range spells {
		le, ri := 0, m
		for le < ri {
			mid := le + (ri-le)/2
			if mid >= 0 && mid < m && int64(ch*potions[mid]) >= success {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		ans[i] = max(0, m-le)
	}
	return ans
}
