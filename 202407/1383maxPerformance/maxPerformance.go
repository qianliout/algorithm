package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3))
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	mod := int(math.Pow10(9)) + 7
	pairs := make([]pair, 0)
	for i := 0; i < n; i++ {
		pairs = append(pairs, pair{sp: speed[i], ef: efficiency[i]})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].ef > pairs[j].ef })

	hq := make(MinHeap, 0)
	sum := 0
	ans := 0
	eff := math.MaxInt / 2
	for i := 0; i < n; i++ {
		eff = min(eff, pairs[i].ef)
		sum += pairs[i].sp
		heap.Push(&hq, pairs[i].sp)
		for hq.Len() > k {
			pop := heap.Pop(&hq).(int)
			sum -= pop
		}
		ans = max(ans, sum*eff)
	}
	return ans % mod
}

type pair struct {
	sp, ef int
}
