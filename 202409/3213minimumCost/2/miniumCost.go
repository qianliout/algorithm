package main

import (
	"fmt"
	"index/suffixarray"
	"math"
)

func main() {
	fmt.Println(minimumCost("abcdef", []string{"abdef", "abc", "d", "def", "ef"}, []int{100, 1, 1, 10, 5}))
}
func minimumCost(target string, words []string, costs []int) int {
	mc := make(map[string]int16)
	for i, c := range words {
		cost, ok := mc[c]
		if !ok || cost > int16(costs[i]) {
			mc[c] = int16(costs[i])
		}
	}
	n := len(target)
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range mc {
		left := sa.Lookup([]byte(w), -1)
		for _, l := range left {
			r := l + len(w)
			from[r] = append(from[r], pair{int16(l), c})
		}
	}
	f := make([]int, n+1)
	inf := math.MaxInt32
	for i := 1; i <= n; i++ {
		// 初值
		f[i] = inf
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.left]+int(p.cost))
		}
	}
	if f[n] == inf {
		return -1
	}
	return f[n]
}

// 直接用 int 会超过内存
type pair struct {
	left int16
	cost int16
}
