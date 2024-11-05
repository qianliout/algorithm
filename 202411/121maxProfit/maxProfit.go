package main

import (
	"math"
)

func main() {

}

func maxProfit(prices []int) int {
	ans := 0
	mi := math.MaxInt
	for i := 0; i < len(prices); i++ {
		ans = max(ans, prices[i]-mi)
		mi = min(mi, prices[i])
	}

	return ans
}
