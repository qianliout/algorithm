package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCostSetTime(0, 1, 4, 9))
}

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	var cal func(mm, ss int) int
	inf := math.MaxInt / 10
	cal = func(mm, ss int) int {
		// 判断数据范围
		if mm < 0 || mm > 99 || ss < 0 || ss > 99 {
			return inf
		}
		digit := []int{mm / 10, mm % 10, ss / 10, ss % 10}
		start := 0
		for start < 4 && digit[start] == 0 {
			start++
		}
		pre := startAt
		cost := 0
		for _, di := range digit[start:] {
			if di != pre {
				pre = di
				cost += moveCost
			}
			cost += pushCost
		}
		return cost
	}

	mm, ss := targetSeconds/60, targetSeconds%60

	return min(cal(mm, ss), cal(mm-1, ss+60))
}
