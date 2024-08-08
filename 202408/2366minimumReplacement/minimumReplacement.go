package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumReplacement([]int{3, 9, 3}))
	fmt.Println(minimumReplacement([]int{12, 9, 7, 6, 17, 19, 21}))
}

func minimumReplacement(nums []int) int64 {
	if len(nums) <= 1 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	m := nums[0]
	ans := 0
	// todo 没有弄懂
	for _, ch := range nums {
		k := (ch - 1) / m
		ans += k
		m = ch / (k + 1)
	}
	return int64(ans)
}
