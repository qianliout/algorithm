package main

import (
	"fmt"
	"sort"
)

func main() {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{1, 1, 5})
	nextPermutation([]int{3, 2, 1})
}

func nextPermutation(nums []int) {
	n := len(nums)
	le := -1
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			le = i - 1
			break
		}
	}
	if le == -1 {
		sort.Ints(nums)
		return
	}
	ri := le + 1
	for i := ri + 1; i < n; i++ {
		if nums[i] > nums[le] && nums[i] < nums[ri] {
			ri = i
		}
	}
	nums[le], nums[ri] = nums[ri], nums[le]
	sort.Ints(nums[le+1:])
	fmt.Println(nums)
}
