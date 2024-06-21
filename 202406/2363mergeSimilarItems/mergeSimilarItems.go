package main

import (
	"sort"
)

func main() {

}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	cnt := make(map[int]int)
	for _, ch := range items1 {
		v, w := ch[0], ch[1]
		cnt[v] += w
	}
	for _, ch := range items2 {
		v, w := ch[0], ch[1]
		cnt[v] += w
	}
	ans := make([][]int, 0)
	for v, w := range cnt {
		ans = append(ans, []int{v, w})
	}

	sort.Slice(ans, func(i, j int) bool { return ans[i][0] < ans[j][0] })
	return ans
}
