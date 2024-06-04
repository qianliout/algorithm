package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumCost([]int{1, 5, 1, 6}))
}

func minimumCost2(nums []int) int {
	nums2 := nums[1:]

	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	return nums[0] + nums2[0] + nums2[1]
}

func minimumCost(nums []int) int {
	se, th := math.MaxInt, math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i] < se {
			th = se
			se = nums[i]
		} else {
			th = min(th, nums[i])
		}
	}
	return nums[0] + se + th
}
