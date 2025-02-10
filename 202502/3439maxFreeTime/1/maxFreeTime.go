package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxFreeTime(5, 1, []int{1, 3}, []int{2, 5}))
}

// 移动之后相对位置不变
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	pp := make([]pair, n)
	for i := 0; i < n; i++ {
		pp[i] = pair{start: startTime[i], end: endTime[i]}
	}
	interval := make([]int, 0)
	sort.Slice(pp, func(i, j int) bool {
		if pp[i].start != pp[j].start {
			return pp[i].start < pp[j].start
		}
		return pp[i].end < pp[j].end
	})

	pre := 0
	for i := 0; i < len(pp); i++ {
		interval = append(interval, pp[i].start-pre)
		pre = pp[i].end
	}
	if eventTime-pre > 0 {
		interval = append(interval, eventTime-pre)
	}
	if interval[0] == 0 {
		interval = interval[1:]
	}

	k++
	wid := 0
	ans := 0
	for i := 0; i < len(interval); i++ {
		if i < k {
			wid += interval[i]
		} else {
			wid = wid + interval[i] - interval[i-k]
		}
		ans = max(ans, wid)
	}
	return ans
}

type pair struct {
	start int
	end   int
}
