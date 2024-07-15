package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(trimMean([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}))
}

func trimMean(arr []int) float64 {
	n := len(arr)
	de := int(math.Ceil(float64(n) * 0.05))
	sort.Ints(arr)
	if n-de*2 <= 0 {
		return 0
	}
	all := 0

	for i := de; i < n-de; i++ {
		all += arr[i]
	}

	return float64(all) / float64(n-2*de)
}
