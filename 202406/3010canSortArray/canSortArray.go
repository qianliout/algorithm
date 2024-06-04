package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	fmt.Println(canSortArray([]int{8, 4, 2, 30, 15}))
	fmt.Println(canSortArray([]int{1, 2, 3, 4, 5}))
}

func canSortArray(nums []int) bool {
	n := len(nums)
	for i := range nums {
		start := i
		ones := bits.OnesCount(uint(nums[i]))
		i++
		for i < n && bits.OnesCount(uint(nums[i])) == ones {
			i++
		}
		slices.Sort(nums[start:i])
	}

	return slices.IsSorted(nums)
}

func bitOneSame(a int) int {
	ans := 0
	for a > 0 {
		ans += a & 1
		a = a >> 1
	}
	return ans
}
