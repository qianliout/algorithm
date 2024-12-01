package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			le, ri := 0, j
			for le < ri {
				mid := le + (ri-le)/2
				if mid >= 0 && mid < j && nums[mid]+nums[j] > nums[i] {
					ri = mid
				} else {
					le = mid + 1
				}
			}
			if le >= 0 && le < j && nums[le]+nums[j] > nums[i] {
				cnt += max(0, j-le)
			}
		}
	}
	return cnt
}
