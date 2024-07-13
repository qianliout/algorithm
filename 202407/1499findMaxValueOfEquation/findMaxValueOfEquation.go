package main

import (
	"math"
)

func main() {

}

func findMaxValueOfEquation(points [][]int, k int) int {
	st := make([]pair, 0)
	ans := math.MinInt / 10

	for _, ch := range points {
		x, y := ch[0], ch[1]
		for len(st) > 0 && x-st[0].X > k {
			st = st[1:]
		}
		if len(st) > 0 {
			ans = max(ans, x+y+st[0].Y_X)
		}
		for len(st) > 0 && st[len(st)-1].Y_X <= y-x {
			st = st[:len(st)-1]
		}

		st = append(st, pair{X: x, Y_X: y - x})
	}
	return ans
}

type pair struct {
	X   int
	Y_X int // y-x
}
