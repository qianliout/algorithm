package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))
	fmt.Println(minimumAverageDifference([]int{4, 2, 0}))
}

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	sum := 0
	pre := make([]int, len(nums)+1)
	for i, ch := range nums {
		pre[i+1] = pre[i] + ch
		sum += ch
	}
	ans := 0

	diff := sum

	for i := 0; i < n-1; i++ {
		di := pre[i+1]/(i+1) - (sum-pre[i+1])/(n-i-1)
		if di < 0 {
			di = -di
		}
		if diff > di {
			ans = i
			diff = di
		}
	}
	if sum/(n) < diff {
		ans = n - 1
	}
	return ans
}
