package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))
	fmt.Println(countSubarrays([]int{2, 5, 1, 4, 3, 6}, 1))
}

func countSubarrays(nums []int, k int) int {
	start := slices.Index(nums, k)
	cnt := make(map[int]int)
	cnt[0] = 1 // 只选一个元素
	a := 0
	for i := start - 1; i >= 0; i-- {
		if nums[i] > nums[start] {
			a--
		} else if nums[i] < nums[start] {
			a++
		}
		cnt[a]++
	}
	// // i=start 的时候 x 是 0，直接加到答案中，这样下面不是大于 k 就是小于 k
	ans := cnt[0] + cnt[-1]
	b := 0
	for i := start + 1; i < len(nums); i++ {
		if nums[i] > nums[start] {
			b++
		} else if nums[i] < nums[start] {
			b--
		}
		ans += cnt[b]
		ans += cnt[b-1]
		// cnt[b]++
	}

	return ans
}
