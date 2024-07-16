package main

import (
	"math"
)

func main() {

}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	var dfs func(i, j int) int

	n := len(aliceValues)
	dfs = func(i, j int) int {
		if i < 0 || j >= n {
			return math.MinInt
		}
		if i == j {
			return aliceValues[i]
		}

	}

}
