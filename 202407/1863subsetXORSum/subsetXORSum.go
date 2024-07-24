package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

func subsetXORSum(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 1; i < 1<<n; i++ {
		ans += cal(nums, i)
	}

	return ans
}

func cal(nums []int, j int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if j>>i&1 == 1 {
			ans = ans ^ nums[i]
		}
	}
	return ans
}
