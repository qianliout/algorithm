package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	inf := math.MaxInt64 / 100
	end := make([]int, n+1)
	for i := range end {
		end[i] = inf
	}
	end[0] = -inf
	ans := make([]int, n)

	mx := 1
	for i := 0; i < n; i++ {
		ch := obstacles[i]
		le, ri := 1, mx+1
		for le < ri {
			mid := le + (ri-le)/2
			if end[mid] > ch {
				ri = mid
			} else {
				le = mid + 1
			}
		}
		ans[i] = le
		end[le] = ch
		mx = max(mx, le)
	}
	return ans
}
