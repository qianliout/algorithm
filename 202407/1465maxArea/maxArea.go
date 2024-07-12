package main

import (
	"math"
	"sort"
)

func main() {

}

func maxArea(h int, w int, hc []int, vc []int) int {
	sort.Ints(hc)
	sort.Ints(vc)
	base := int(math.Pow10(9)) + 7
	return getMax(h, hc) * getMax(w, vc) % base

}
func getMax(size int, hc []int) int {
	n := len(hc)

	res := max(hc[0], size-hc[n-1])
	for i := 1; i < n; i++ {
		res = max(hc[i]-hc[i-1], res)
	}
	return res
}
