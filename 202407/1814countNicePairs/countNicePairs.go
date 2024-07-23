package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countNicePairs([]int{42, 11, 1, 97}))
	fmt.Println(countNicePairs([]int{13, 10, 35, 24, 76}))
}

func countNicePairs(nums []int) int {
	n := len(nums)
	nums2 := make([]int, n)
	for i := range nums2 {
		nums2[i] = nums[i] - rev(nums[i])
	}
	cnt := make(map[int]int)
	for _, ch := range nums2 {
		cnt[ch]++
	}
	ans := 0
	mod := int(math.Pow10(9)) + 7
	for _, v := range cnt {
		ans += (v * (v - 1)) / 2
	}

	return ans % mod
}

func rev(n int) int {
	ans := 0
	for n > 0 {
		ans = ans*10 + n%10
		n = n / 10
	}
	return ans
}
