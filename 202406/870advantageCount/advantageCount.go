package main

import (
	"sort"
)

func main() {

}

func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	pairs := make([]pair, len(nums2))
	for i := range nums2 {
		pairs[i] = pair{va: nums2[i], idx: i}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].va < pairs[j].va })
	n := len(nums1)
	ans := make([]int, n)
	le, ri := 0, n-1

	for _, x := range nums1 {
		if x > pairs[le].va {
			ans[pairs[le].idx] = x
			le++
		} else {
			ans[pairs[ri].idx] = x
			ri--
		}
	}
	return ans
}

type pair struct {
	va  int
	idx int
}
