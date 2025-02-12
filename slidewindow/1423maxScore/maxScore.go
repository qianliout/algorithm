package main

import (
	"math"
)

func main() {

}

func maxScore(cardPoints []int, k int) int {
	sum := 0
	for _, ch := range cardPoints {
		sum += ch
	}
	win, mi := 0, math.MaxInt64
	le, ri, n := 0, 0, len(cardPoints)
	for le <= ri && ri < n {
		win += cardPoints[ri]
		ri++
		if ri-le == n-k {
			mi = min(mi, win)
		}
		if ri-le >= n-k {
			win -= cardPoints[le]
			le++
		}
	}
	if mi == math.MaxInt64 {
		return sum
	}
	return sum - mi
}
