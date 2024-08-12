package main

import (
	"slices"
)

func main() {

}

func destroyTargets(nums []int, space int) int {
	cnt := make(map[int][]int)
	for _, ch := range nums {
		cnt[ch%space] = append(cnt[ch%space], ch)
	}
	mx, ans := 0, 0
	for _, ch := range cnt {
		m := len(ch)
		mi := slices.Min(ch)
		if m > mx || (m == mx && mi < ans) {
			mx = max(m, mx)
			ans = mi
		}
	}
	return ans
}
