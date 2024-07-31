package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minOperations([]int{1, 10, 100, 1000}))
	fmt.Println(minOperations([]int{1, 2, 3, 5, 6}))
	fmt.Println(minOperations([]int{4, 2, 5, 3}))
	fmt.Println(minOperations([]int{8, 5, 9, 9, 8, 4}))
	fmt.Println(minOperations([]int{8, 10, 16, 18, 10, 10, 16, 13, 13, 16}))
}

func minOperations(nums []int) int {
	n, mx := len(nums), 1
	nums = reset(nums)
	sort.Ints(nums)
	le, ri := 0, 0

	for ri < len(nums) {
		for le < ri && nums[ri]-nums[le] > n-1 {
			le++
		}
		mx = max(mx, ri-le+1)
		ri++
	}

	return n - mx
}

// 去重
func reset(nums []int) []int {
	cnt := make(map[int]int)
	ans := make([]int, 0)
	for _, ch := range nums {
		if cnt[ch] == 0 {
			ans = append(ans, ch)
		}
		cnt[ch]++
	}
	return ans
}
