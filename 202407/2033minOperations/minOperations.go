package main

import (
	"sort"
)

func main() {

}

func minOperations(grid [][]int, x int) int {
	nums := make([]int, 0)
	for _, row := range grid {
		for _, ch := range row {
			nums = append(nums, ch)
		}
	}
	sort.Ints(nums)

	mid := nums[len(nums)/2]
	ans := 0

	for _, ch := range nums {
		if abs(ch-nums[0])%x != 0 {
			return -1
		}
		ans += abs(ch-mid) / x
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
