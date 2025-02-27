package main

import (
	"sort"
)

func main() {

}

func nextPermutation(nums []int) {
	pre, next, n := -1, -1, len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pre = i - 1
			next = i
			break
		}
	}
	if pre == -1 {
		sort.Ints(nums)
		return
	}
	sort.Ints(nums[next:])
	for j := next; j < n; j++ {
		if nums[j] > nums[pre] {
			nums[j], nums[pre] = nums[pre], nums[j]
			break
		}
	}
}
