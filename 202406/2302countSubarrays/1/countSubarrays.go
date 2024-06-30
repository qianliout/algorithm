package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubarrays([]int{2, 1, 4, 3, 5}, 10))
}

// 时间复杂度太高，会超时
func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	var ans int64
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if int64(sum[j+1]-sum[i])*int64(j-i+1) < k {
				ans++
			}

		}
	}
	return ans
}
