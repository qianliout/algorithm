package main

import (
	"sort"
)

func main() {

}

func maximumScore(a int, b int, c int) int {
	nums := []int{a, b, c}
	sort.Ints(nums)
	a, b, c = nums[0], nums[1], nums[2]
	if a+b > c {
		return a + b
	}
	return (a + b + c) >> 1
}
