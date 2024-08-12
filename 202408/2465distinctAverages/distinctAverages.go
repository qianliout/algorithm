package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distinctAverages([]int{4, 1, 4, 0, 3, 5}))
}

func distinctAverages(nums []int) int {
	cnt := make(map[int]int)
	sort.Ints(nums)
	for len(nums) > 0 {
		a, b := nums[0], nums[len(nums)-1]
		cnt[a+b]++
		nums = nums[1 : len(nums)-1]
	}

	return len(cnt)
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}

type pair struct {
	a, b int
}
