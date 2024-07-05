package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	// fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	type pair struct{ start, end, profit int }
	work := make([]pair, n)
	for i := range startTime {
		work[i] = pair{start: startTime[i], end: endTime[i], profit: profit[i]}
	}
	sort.Slice(work, func(i, j int) bool {
		if work[i].start < work[j].start {
			return true
		} else if work[i].start > work[j].start {
			return false
		} else {
			return work[i].end <= work[j].end
		}
	})
	var dfs func(i int) int
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i < 0 || i >= n {
			return 0
		}
		if mem[i] >= 0 {
			return mem[i]
		}
		do := work[i].profit // 做
		for j := i + 1; j < n; j++ {
			if work[j].start >= work[i].end {
				do += dfs(j)
				break
			}
		}
		not := dfs(i + 1)
		mem[i] = max(do, not)
		return mem[i]
	}

	res := dfs(0)
	return res
}
