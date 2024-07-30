package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeArray([]int{1, 2, 3, 4, 5}))
}

// 题目中说互不相同
func rearrangeArray(nums []int) []int {
	sort.Ints(nums)

	n := len(nums)
	ans := make([]int, n)
	i := 0
	le, ri := 0, n-1
	for le <= ri && i < n {
		ans[i] = nums[le]
		if i+1 < n {
			ans[i+1] = nums[ri]
		}
		i += 2
		le++
		ri--
	}

	return ans
}
