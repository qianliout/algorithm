package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubsequence([]int{1, 1, 5, 4, 5}, 3))
	fmt.Println(lengthOfLongestSubsequence([]int{4, 1, 3, 2, 1, 5}, 7))
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	f := make([]int, target+10)
	inf := len(nums) * 2
	for i := range f {
		f[i] = -inf
	}
	f[0] = 0
	// 0-1背包问题
	s := 0
	for _, ch := range nums {
		s = min(s+ch, target)
		for j := s; j >= ch; j-- {
			f[j] = max(f[j], f[j-ch]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}
