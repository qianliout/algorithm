package main

import (
	"math"
)

func main() {

}

func maximumEnergy(energy []int, k int) int {
	ans := math.MinInt
	n := len(energy)
	for i := n - 1; i >= n-k; i-- {
		cnt := 0
		for j := i; j >= 0; j = j - k {
			cnt += energy[j]
			ans = max(ans, cnt)
		}
	}
	return ans
}
