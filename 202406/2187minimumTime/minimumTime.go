package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	fmt.Println(minimumTime([]int{2}, 1))
}

func minimumTime(nums []int, totalTrips int) int64 {
	mi := slices.Min(nums)
	le, ri := 0, mi*totalTrips+1
	for le < ri {
		mid := le + (ri-le)/2
		// 相当于左端点
		if mid > 0 && mid < ri && f(nums, mid) >= totalTrips {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return int64(le)
}

// 表示在这个时间点后，能完成的
func f(nums []int, t int) int {
	ans := 0
	for _, ch := range nums {
		ans += t / ch
	}
	return ans
}
