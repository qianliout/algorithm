package main

import (
	"sort"
)

func main() {}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		}
		return intervals[i][1] <= intervals[j][1]
	})

	ans := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		inv := intervals[i]
		if end < inv[0] {
			ans = append(ans, []int{start, end})
			start, end = inv[0], inv[1]
		} else {
			end = max(end, inv[1])
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}
