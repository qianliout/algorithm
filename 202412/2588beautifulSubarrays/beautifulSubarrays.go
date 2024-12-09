package main

import (
	"fmt"
)

func main() {
	fmt.Println(beautifulSubarrays([]int{4, 3, 1, 2, 4}))
}

// 会超时
func beautifulSubarrays1(nums []int) int64 {
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] ^ nums[i]
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if pre[i]-pre[j] == 0 {
				ans++
			}
		}
	}
	return int64(ans)
}

func beautifulSubarrays(nums []int) int64 {
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] ^ nums[i]
	}
	ans := 0
	cnt := make(map[int]int)
	for _, ch := range pre {
		ans += cnt[ch]
		cnt[ch]++
	}

	return int64(ans)
}
