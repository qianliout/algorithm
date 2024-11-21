package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimizeSum([]int{59, 27, 9, 81, 33}))
}

func minimizeSum(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)

	n := len(nums)
	ans := math.MaxInt64
	for i := 0; i < 3; i++ {
		a := nums[i]
		b := nums[i+n-3]
		ans = min(ans, b-a)
	}
	return ans
}
