package main

import (
	"fmt"
)

func main() {
}

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	if len(nums) == 0 {
		return ans
	}
	n := len(nums)
	start, end := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i]-end == 1 {
			end++
		} else {
			if start == end {
				ans = append(ans, fmt.Sprintf("%d", start))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", start, end))
			}
			start = nums[i]
			end = nums[i]
		}
	}
	if start == end {
		ans = append(ans, fmt.Sprintf("%d", start))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", start, end))
	}
	return ans
}
