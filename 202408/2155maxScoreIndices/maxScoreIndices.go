package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxScoreIndices([]int{0, 0, 1, 0}))

}

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	zero, one := make([]int, n+1), make([]int, n+1)

	for i := 0; i < n; i++ {
		zero[i+1] = zero[i]
		if nums[i] == 0 {
			zero[i+1] += 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		one[i] = one[i+1]
		if nums[i] == 1 {
			one[i] += 1
		}
	}
	ans := make([]int, 0)
	mx := 0
	for i := 0; i <= n; i++ {
		a := zero[i] + one[i]
		if a > mx {
			mx = a
			ans = []int{i}
		} else if a == mx {
			ans = append(ans, i)
		}
	}
	return ans
}
