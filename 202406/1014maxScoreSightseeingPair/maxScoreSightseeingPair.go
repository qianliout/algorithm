package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	res := 0
	preMax := values[0] + 0
	for i := 1; i < n; i++ {
		res = max(res, preMax+(values[i]-i))
		preMax = max(preMax, values[i]+i)
	}
	return res
}
