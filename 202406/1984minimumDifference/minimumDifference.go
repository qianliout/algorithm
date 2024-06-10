package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	wind := make([]int, 0)
	for ri := 0; ri < len(nums); ri++ {
		wind = append(wind, nums[ri])
		for len(wind) > k {
			wind = wind[1:]
		}
		if len(wind) == k {
			ans = min(ans, cac(wind))
		}
	}

	return ans
}

func cac(nums []int) int {
	ma, mi := nums[0], nums[0]
	for _, ch := range nums {
		ma = max(ma, ch)
		mi = min(mi, ch)
	}
	return ma - mi
}
