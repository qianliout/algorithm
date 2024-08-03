package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSubsequence([]int{2, 1, 3, 3}, 2))
}

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	ids := make([]pair, n)
	for i := range ids {
		ids[i] = pair{i, nums[i]}
	}
	sort.SliceStable(ids, func(i, j int) bool { return ids[i].val >= ids[j].val })
	ids = ids[:k]

	sort.SliceStable(ids, func(i, j int) bool { return ids[i].id < ids[j].id })
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, ids[i].val)
	}
	return ans
}

type pair struct {
	id  int
	val int
}
