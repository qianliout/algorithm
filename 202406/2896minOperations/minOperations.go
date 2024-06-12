package main

import (
	"sort"
)

func main() {

}

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	ans := make([]int64, len(queries))
	for i, ch := range queries {
		// 查的右端点,这样才能得到正确的答案
		j := sort.SearchInts(nums, ch)
		ll := (j * ch) - sum[j]
		rr := sum[n] - sum[j] - (n-j)*ch

		ans[i] = int64(ll + rr)
	}
	return ans
}
