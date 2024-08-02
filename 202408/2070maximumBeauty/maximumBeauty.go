package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))
}

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool { return items[i][0] < items[j][0] })
	n := len(queries)
	ans := make([]int, n)
	que := make([]pair, 0)
	for i, ch := range queries {
		que = append(que, pair{i, ch})
	}
	sort.Slice(que, func(i, j int) bool { return que[i].val < que[j].val })

	mx, j := 0, 0
	for _, ch := range que {
		for j < len(items) && items[j][0] <= ch.val {
			mx = max(mx, items[j][1])
			j++
		}
		ans[ch.idx] = mx
	}
	return ans
}

type pair struct {
	idx int
	val int
}
