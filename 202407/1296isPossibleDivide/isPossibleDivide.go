package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
}

// 没有要求子数组或子序列
func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}

	cnt := make(map[int]int)
	mi := nums[0]
	for _, ch := range nums {
		cnt[ch]++
		mi = min(mi, ch)
	}

	used := 0
	for used < n {
		preMi := findMin(cnt)
		for start := preMi; start < preMi+k; start++ {
			if cnt[start] <= 0 {
				return false
			}
			cnt[start]--
			used++
		}
	}

	return true
}

func findMin(cnt map[int]int) int {
	mi := math.MaxInt
	for k, v := range cnt {
		if v > 0 {
			mi = min(k, mi)
		}
	}
	return mi
}
