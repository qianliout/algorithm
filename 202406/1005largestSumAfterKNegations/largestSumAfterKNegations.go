package main

import (
	"sort"
)

func main() {

}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	pos := 0
	hasZ := false
	for i := 0; i < n; i++ {
		if pos >= k {
			break
		}
		if nums[i] == 0 {
			hasZ = true
			break
		}
		if nums[i] > 0 {
			break
		}
		if nums[i] < 0 {
			nums[i] = -nums[i]
			pos++
		}
	}
	sort.Ints(nums)
	sum := 0
	for _, ch := range nums {
		sum += ch
	}

	if pos >= k || hasZ || (k-pos)%2 == 0 {
		return sum
	}
	fir := nums[0]
	if fir < 0 {
		sum += -2 * fir
	} else {
		sum -= 2 * fir
	}
	return sum
}
