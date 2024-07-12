package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := make(map[int]int)
	for _, ch := range arr {
		cnt[ch]++
	}
	ans := make([]pair, 0)
	for ke, v := range cnt {
		ans = append(ans, pair{va: ke, cnt: v})
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i].cnt < ans[j].cnt })

	start := 0
	for start < len(ans) && k > 0 {
		if ans[start].cnt <= k {
			k -= ans[start].cnt
			start++
		} else {
			break
		}
	}
	fmt.Println(ans, start)
	return len(ans) - start
}

type pair struct {
	va  int
	cnt int
}
