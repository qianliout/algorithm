package main

import (
	"math"
)

func main() {

}

func maxMatrixSum(matrix [][]int) int64 {
	cnt := 0
	all := 0
	mi := math.MaxInt64 / 100

	for i := range matrix {
		for _, ch := range matrix[i] {
			all += abs(ch)
			mi = min(mi, abs(ch))
			if ch < 0 {
				cnt++
			}
		}
	}
	if cnt&1 == 0 {
		return int64(all)
	}
	return int64(all - 2*mi)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
