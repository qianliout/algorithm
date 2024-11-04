package main

import (
	"slices"
	"sort"
)

func main() {

}

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	sort.Ints(hFences)
	sort.Ints(vFences)
	h := make(map[int]int)
	v := make(map[int]int)
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			h[hFences[j]-hFences[i]]++
		}
	}
	for i := 0; i < len(vFences); i++ {
		for j := i + 1; j < len(vFences); j++ {
			v[vFences[j]-vFences[i]]++
		}
	}
	res := mix(h, v)
	if len(res) == 0 {
		return -1
	}
	a := slices.Max(res)
	return a * a % 1_000_000_007
}

func mix(a, b map[int]int) []int {
	ans := make([]int, 0)
	for k := range a {
		if b[k] > 0 {
			ans = append(ans, k)
		}
	}
	return ans
}
