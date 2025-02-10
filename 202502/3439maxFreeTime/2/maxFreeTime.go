package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxFreeTime(5, 1, []int{1, 3}, []int{2, 5}))
}

// 移动之后相对位置不变
// 会议不重复
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	interval := make([]int, 0)

	pre := 0
	for i := 0; i < n; i++ {
		interval = append(interval, startTime[i]-pre)
		pre = endTime[i]
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
