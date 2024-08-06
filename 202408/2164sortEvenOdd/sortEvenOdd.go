package main

import (
	"sort"
)

func main() {

}

func sortEvenOdd(nums []int) []int {
	n := len(nums)
	even, odd := make([]int, 0), make([]int, 0)
	for i := 0; i < n; i = i + 2 {
		even = append(even, nums[i])
		if i+1 < n {
			odd = append(odd, nums[i+1])
		}
	}
	sort.SliceStable(even, func(i, j int) bool { return even[i] < even[j] })
	sort.SliceStable(odd, func(i, j int) bool { return odd[i] > odd[j] })
	ans := make([]int, 0)
	for i := 0; i < len(even); i++ {
		ans = append(ans, even[i])
		if i < len(odd) {
			ans = append(ans, odd[i])
		}
	}
	return ans
}
