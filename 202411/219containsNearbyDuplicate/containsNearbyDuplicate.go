package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	cnt := make(map[int]int)
	for i, ch := range nums {
		if pre, ok := cnt[ch]; ok {
			if i != pre && abs(pre-i) <= k {
				return true
			}
		}
		cnt[ch] = i
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
