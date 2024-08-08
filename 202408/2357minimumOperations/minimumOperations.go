package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
}

func minimumOperations(nums []int) int {
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	ans := 0
	for sum > 0 {
		n := findMin(nums)
		for i := range nums {
			if nums[i] == 0 {
				continue
			}
			nums[i] -= n
			sum -= n
		}
		ans++
	}
	return ans
}

func findMin(nums []int) int {
	ans := math.MaxInt
	for _, ch := range nums {
		if ch != 0 {
			ans = min(ans, ch)
		}
	}
	return ans
}
