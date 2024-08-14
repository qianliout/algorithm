package main

import (
	"math"
	"sort"
)

func main() {

}

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
	g := make([]sort.IntSlice, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		if vals[y] > 0 {
			g[x] = append(g[x], vals[y])
		}
		if vals[x] > 0 {
			g[y] = append(g[y], vals[x])
		}

	}
	ans := math.MinInt32
	for i, a := range g {
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		if k <= len(a) {
			a = a[:k]
		}
		sum := vals[i]
		for _, ch := range a {
			sum += ch
		}
		ans = max(ans, sum)
	}
	return ans
}
