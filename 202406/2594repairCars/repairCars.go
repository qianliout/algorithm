package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(repairCars([]int{4, 2, 3, 1}, 10))
}

func repairCars(ranks []int, cars int) int64 {
	mx := slices.Max(ranks) * cars * cars
	le, ri := 0, mx

	for le < ri {
		mid := le + (ri-le)/2
		if f(ranks, mid) >= cars {
			ri = mid
		} else {
			le = mid + 1
		}
	}

	return int64(le)
}

// t的时间内，能修好多少个车
func f(ranks []int, t int) int {
	ans := 0
	for _, ch := range ranks {
		n := int(math.Sqrt(float64(t / ch)))
		ans += n
	}
	return ans
}
