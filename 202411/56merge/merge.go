package main

import (
	"sort"
)

func main() {

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	ans := make([][]int, 0)
	n := len(intervals)
	if n == 0 {
		return ans
	}
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		x, y := intervals[i][0], intervals[i][1]
		if x <= end {
			end = max(end, y)
		} else {
			ans = append(ans, []int{start, end})
			start, end = x, y
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}
