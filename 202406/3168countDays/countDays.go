package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
	fmt.Println(countDays(5, [][]int{{2, 4}, {1, 3}}))
}

func countDays1(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] < meetings[j][0] {
			return true
		} else if meetings[i][0] > meetings[j][0] {
			return false
		}
		return meetings[i][1] < meetings[j][1]
	})
	start, end := meetings[0][0], meetings[0][1]
	res := start - 1
	for i := 1; i < len(meetings); i++ {
		me := meetings[i]
		if me[0] <= end {
			end = max(me[1], end)
		} else {
			res += me[0] - end - 1
			start, end = me[0], me[1]
		}
	}
	res += max(0, days-end)
	return res
}

// 总天数减去开会的天数
func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] < meetings[j][0] {
			return true
		} else if meetings[i][0] > meetings[j][0] {
			return false
		}
		return meetings[i][1] < meetings[j][1]
	})
	start, end := meetings[0][0], meetings[0][1]
	res := 0
	for i := 1; i < len(meetings); i++ {
		me := meetings[i]
		if me[0] <= end {
			end = max(me[1], end)
		} else {
			res += end - start + 1
			start, end = me[0], me[1]
		}
	}

	res += end - start + 1
	return days - res
}
