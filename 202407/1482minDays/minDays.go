package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
}

func minDays(bloomDay []int, m int, k int) int {
	mx := slices.Max(bloomDay) + 1
	le, ri := 1, mx
	for le < ri {
		mid := le + (ri-le)/2
		if le > 0 && le < mx && check(bloomDay, k, mid) >= m {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le >= 1 && le < mx {
		return le
	}
	return -1
}

func check(nums []int, k, mid int) int {
	fl := make([]int, len(nums))
	for i, ch := range nums {
		if ch <= mid {
			fl[i] = ch
		}
	}
	// 开始统计
	start := 0
	ans := 0
	for start < len(nums) {
		cnt, j := 0, start
		for j < len(nums) {
			if fl[j] <= 0 {
				break
			}
			cnt++
			if cnt >= k {
				ans++
				break
			}
			j++
		}
		start = j + 1
	}
	return ans
}
