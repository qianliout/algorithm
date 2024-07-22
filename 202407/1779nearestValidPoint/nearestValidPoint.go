package main

import (
	"math"
)

func main() {

}

func nearestValidPoint(x int, y int, points [][]int) int {
	ans := -1
	sub := math.MaxInt / 10
	for i, ch := range points {
		a, b := ch[0], ch[1]
		if a != x && b != y {
			continue
		}
		c := abs(a-x) + abs(b-y)
		if c < sub {
			ans = i
		}
		sub = min(sub, c)
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
