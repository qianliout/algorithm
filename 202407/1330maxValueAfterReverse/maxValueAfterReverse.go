package main

import (
	"math"
)

func main() {

}

func maxValueAfterReverse(nums []int) int {
	base := 0
	d := 0
	n := len(nums)
	mx := math.MinInt / 10 //
	mi := math.MaxInt / 10

	for i := 1; i < n; i++ {
		a, b := nums[i-1], nums[i]
		base += abs(a - b)
		mx = max(mx, min(a, b))
		mi = min(mi, max(a, b))
		d = max(d, max(abs(nums[0]-b)-abs(a-b), // i=0
			abs(nums[n-1]-a)-abs(a-b))) // j=n-1
	}
	return base + max(d, 2*(mx-mi))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
