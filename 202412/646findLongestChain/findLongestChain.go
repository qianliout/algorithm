package main

import (
	"sort"
)

func main() {

}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] != pairs[j][0] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})

	n := len(pairs)
	f := make([]int, n)
	mx := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := i - 1; j >= 0; j-- {
			if pairs[j][1] < pairs[i][0] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		mx = max(mx, f[i])
	}
	return mx
}
