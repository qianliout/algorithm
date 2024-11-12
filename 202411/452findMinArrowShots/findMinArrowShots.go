package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
}

func findMinArrowShots1(intervals [][]int) int {
	// 以终点排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	if n == 0 {
		return 0
	}
	ans := 1 // 上面判断了 n==0 的情况，这里把初值设置成1是这个题目的关键
	end := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] > end {
			end = intervals[i][1]
			ans++
		}
	}
	return ans
}

func findMinArrowShots(intervals [][]int) int {
	// 以起点排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	n := len(intervals)
	if n == 0 {
		return 0
	}
	ans := 1 // 上面判断了 n==0 的情况，这里把初值设置成1是这个题目的关键
	end := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] > end {
			end = intervals[i][1]
			ans++
		} else {
			end = min(end, intervals[i][1])
		}
	}
	return ans
}
