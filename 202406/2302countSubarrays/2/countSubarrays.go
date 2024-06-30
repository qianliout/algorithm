package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubarrays([]int{2, 1, 4, 3, 5}, 10))
}

// 滑动窗口
func countSubarrays1(nums []int, k int64) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	var ans int64
	wind := make([]int, 0)
	le, ri := 0, 0
	for le <= ri && ri < n {
		wind = append(wind, ri)
		ri++
		for le <= ri && int64(sum[ri]-sum[le])*int64(ri-le) >= k {
			le++
		}
		if int64(sum[ri]-sum[le])*int64(ri-le) < k {
			ans += int64(ri - le)
		}
	}

	return ans
}

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	var sum int64
	var ans int64
	wind := make([]int, 0)
	le, ri := 0, 0
	for le <= ri && ri < n {
		sum += int64(nums[ri])
		wind = append(wind, ri)
		for le <= ri && sum*int64(ri-le+1) >= k {
			sum -= int64(nums[le])
			le++
		}
		if sum*int64(ri-le+1) < k {
			ans += int64(ri - le + 1)
		}
		ri++
	}
	return ans
}
