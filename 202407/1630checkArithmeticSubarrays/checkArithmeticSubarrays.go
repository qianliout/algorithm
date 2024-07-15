package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(checkArithmeticSubarrays([]int{-3, -6, -8, -4, -2, -8, -6, 0, 0, 0, 0}, []int{5, 4, 3, 2, 4, 7, 6, 1, 7}, []int{6, 5, 6, 3, 7, 10, 7, 4, 10}))
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = check(nums, l[i], r[i])
	}
	return ans
}

func check(nums []int, l, r int) bool {
	mx := math.MinInt / 10
	mi := math.MaxInt / 10

	set := make(map[int]bool)
	for i := l; i <= r; i++ {
		mx = max(mx, nums[i])
		mi = min(mi, nums[i])
		set[nums[i]] = true
	}

	d := (mx - mi) / (r - l) // 这里有 r-l+1个元素值
	if d*(r-l) != mx-mi {
		return false
	}
	if d == 0 {
		return true
	}
	for i := mi; i <= mx; i += d {
		if !set[i] {
			return false
		}
	}
	return true
}
