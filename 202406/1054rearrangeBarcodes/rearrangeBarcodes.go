package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3}))
	fmt.Println(rearrangeBarcodes([]int{2, 2, 9, 9, 9, 9, 10, 10, 10, 9}))
}

func rearrangeBarcodes(barcodes []int) []int {
	mx := slices.Max(barcodes)
	cnt := make([]int, mx+1)

	for _, ch := range barcodes {
		cnt[ch]++
	}
	idx := make([]int, mx+1)
	for i := range cnt {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})

	ans := make([]int, len(barcodes))
	n := len(barcodes)
	for i := 0; i < n; i = i + 2 {
		for j := 0; j < len(cnt); j++ {
			if cnt[idx[j]] == 0 {
				continue
			}
			ans[i] = idx[j]
			cnt[idx[j]]--
			break
		}
	}
	for i := 1; i < n; i = i + 2 {
		for j := 0; j < len(cnt); j++ {
			if cnt[idx[j]] == 0 {
				continue
			}
			ans[i] = idx[j]
			cnt[idx[j]]--
			break
		}
	}

	return ans
}
