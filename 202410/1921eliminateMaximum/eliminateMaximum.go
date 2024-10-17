package main

import (
	"sort"
)

func main() {

}

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	pairs := make([]pair, n)
	for i := range dist {
		pairs[i] = pair{dis: dist[i], last: (dist[i] - 1) / speed[i], idx: i}
	}
	// 最先消灭最先到达的怪物
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].last < pairs[j].last })
	for i := range pairs {
		// 在i 时间还不能消灭，说明就没有机会消灭了
		if i > pairs[i].last {
			return i // 因为小标是从0开始的，所以这里不需要减一
		}
	}
	return n
}

type pair struct {
	idx  int
	dis  int
	last int // 最迟消灭的时间
}
