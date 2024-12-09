package main

import (
	"sort"
)

func main() {

}

func findMinimumTime(tasks [][]int) int {
	tt := make([]task, 0)
	for _, ch := range tasks {
		ss := task{ch[0], ch[1], ch[2]}
		tt = append(tt, ss)
	}
	sort.Slice(tt, func(i, j int) bool {
		return tt[i].end < tt[j].end
	})
	n := len(tasks)
	mxEnd := tt[n-1].end
	run := make([]bool, mxEnd+1)
	ans := 0
	for _, ta := range tt {
		st, end, dur := ta.start, ta.end, ta.dur
		// 如果有其他任务已经执行过，那这个任务可以并行执行
		for i := st; i <= end; i++ {
			if run[i] {
				dur--
			}
		}
		// 如果没有执行完，就得向前继续找时间执行
		// 就尽量从后向前执行
		for i := end; dur > 0; i-- {
			if !run[i] {
				run[i] = true
				dur--
				ans++
			}
		}
	}
	return ans
}

type task struct {
	start, end, dur int
}
