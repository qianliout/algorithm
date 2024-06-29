package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 2, 2, 3, 3, 1, 1}))
	fmt.Println(rearrangeBarcodes([]int{2, 2, 9, 9, 9, 9, 10, 10, 10, 9}))
}

func rearrangeBarcodes(barcodes []int) []int {
	mx := slices.Max(barcodes)
	cnt := make([]int, mx+1)

	for _, ch := range barcodes {
		cnt[ch]++
	}
	sort.Slice(barcodes, func(i, j int) bool {
		a, b := barcodes[i], barcodes[j]
		if cnt[a] == cnt[b] {
			return a < b
		}
		return cnt[a] > cnt[b]
	})
	ans := make([]int, len(barcodes))
	n := len(barcodes)
	// 这样写不好理解
	// for k, j := 0, 0; k < 2; k++ {
	// 	for i := k; i < n; i, j = i+2, j+1 {
	// 		ans[i] = barcodes[j]
	// 	}
	// }

	// 元素依次填入答案数组的 0,2,4,⋯ 等偶数下标位置，然后将剩余元素依次填入答案数组的 1,3,5,⋯ 等奇数下标位置即可。
	j := 0
	for i := 0; i < n; i = i + 2 {
		ans[i] = barcodes[j]
		j++
	}
	for i := 1; i < n; i = i + 2 {
		ans[i] = barcodes[j]
		j++
	}

	return ans
}
