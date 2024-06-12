package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
}

func smallestDivisor(nums []int, threshold int) int {
	sort.Ints(nums)
	le, ri := 1, nums[len(nums)-1]
	for le < ri {
		mid := le + (ri-le)/2
		// 求左端点
		if mid >= 1 && mid < ri && f(nums, mid) <= threshold {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func f(nums []int, n int) int {
	ans := 0
	for _, ch := range nums {
		ans += (ch + n - 1) / n
	}
	return ans
}
