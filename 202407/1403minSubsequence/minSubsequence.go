package main

import (
	"sort"
)

func main() {

}

func minSubsequence(nums []int) []int {
	n := len(nums)
	pairs := make([]pair, n)
	sum := 0
	for i, ch := range nums {
		pairs = append(pairs, pair{idx: i, value: ch})
		sum += ch
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].value > pairs[j].value })
	ans := make([]int, 0)
	cnt := 0
	for {
		ans = append(ans, pairs[i].value)
		cnt += pairs[i].value

	}

}

type pair struct {
	idx, value int
}
