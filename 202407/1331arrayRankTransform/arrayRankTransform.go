package main

import (
	"sort"
)

func main() {

}

func arrayRankTransform(arr []int) []int {
	cnt := make(map[int]bool)
	ans := make([]int, 0)
	for _, ch := range arr {
		if cnt[ch] {
			continue
		}
		cnt[ch] = true
		ans = append(ans, ch)
	}
	sort.Ints(ans)
	cnt2 := make(map[int]int)
	for i, ch := range ans {
		cnt2[ch] = i + 1 // 从1 开始
	}
	for i := range arr {
		arr[i] = cnt2[arr[i]]
	}
	return arr

}
