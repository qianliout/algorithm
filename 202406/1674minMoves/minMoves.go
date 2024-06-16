package main

import (
	"fmt"
)

func main() {
	fmt.Println(minMoves([]int{1, 2, 1, 2}, 2))
	fmt.Println(minMoves([]int{1, 2, 4, 3}, 4))
}

func minMoves(nums []int, limit int) int {
	d := make([]int, 2*limit+2)

	n := len(nums)
	for i := 0; i < n/2; i++ {
		lo := 1 + min(nums[i], nums[n-i-1])
		up := limit + max(nums[i], nums[n-i-1])
		sum := nums[i] + nums[n-i-1]

		d[lo]--
		d[sum]--
		d[up+1]++
		d[sum+1]++
	}
	ans := n
	all := n
	// for i := 0; i <= 2*limit; i++ {
	for i := range d {
		all += d[i]
		ans = min(ans, all)
	}
	return ans
}
