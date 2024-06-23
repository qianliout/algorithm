package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))
	fmt.Println(canReorderDoubled([]int{2, 4, 0, 0, 8, 1}))
}

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	cnt := make(map[int]int)
	for _, ch := range arr {
		cnt[ch]++
	}
	for _, ch := range arr {
		if cnt[ch] <= 0 {
			continue
		}
		cnt[ch]--

		if cnt[ch*2] > 0 {
			cnt[ch*2]--
			continue
		}

		if ch%2 == 0 && cnt[ch/2] > 0 {
			cnt[ch/2]--
			continue
		}
		return false
	}
	return true
}
