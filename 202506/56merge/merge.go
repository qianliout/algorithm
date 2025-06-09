package main

import (
	"sort"
)

func main() {

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	n := len(intervals)
	ans := make([][]int, 0)
	for i := 1; i < n; i++ {
		fir, sec := intervals[i][0], intervals[i][1]
		if fir > end {
			ans = append(ans, []int{start, end})
			start, end = fir, sec
		} else {
			end = max(end, sec)
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}
