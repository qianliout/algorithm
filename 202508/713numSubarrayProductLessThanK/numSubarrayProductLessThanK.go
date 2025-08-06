package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	wind := make([]int, 0)
	cnt := 1
	ans := 0
	for i := 0; i < n; i++ {
		wind = append(wind, i)
		cnt = cnt * nums[i]
		for cnt >= k && len(wind) > 0 {
			cnt = cnt / nums[wind[0]]
			wind = wind[1:]
		}
		if len(wind) > 0 {
			ans += len(wind)
		}
	}

	return ans
}
