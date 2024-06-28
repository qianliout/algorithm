package main

import (
	"fmt"
	"sort"
)

func main() {

}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	ans := make([][]int, 0)
	for i := 0; i <= rows; i++ {
		for j := 0; j <= cols; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		a := abs(ans[i][0]-rCenter) + abs(ans[i][1]-cCenter)
		b := abs(ans[j][0]-rCenter) + abs(ans[j][1]-cCenter)
		return a < b
	})
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBoomerang(ps [][]int) bool {
	a, b, c := ps[0], ps[1], ps[2]
	cnt := make(map[string]int)
	for i := 0; i < len(ps); i++ {
		key := fmt.Sprintf("%d-%d", ps[i][0], ps[i][1])
		if cnt[key] > 0 {
			return false
		}
		cnt[key]++
	}

	return (b[0]-a[0])*(c[1]-a[1]) != (c[0]-a[0])*(b[1]-a[1])
}
