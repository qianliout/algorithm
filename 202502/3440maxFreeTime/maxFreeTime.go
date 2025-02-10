package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxFreeTime(10, []int{0, 3, 7, 9}, []int{1, 4, 8, 10}))
}

// 至多调整一个会议
// 移动之后相对位置可以变动
func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	// 获取i 左边的间隔
	get := func(i int) int {
		if i == 0 {
			return startTime[0]
		} else if i == n {
			return eventTime - endTime[i-1]
		}
		return startTime[i] - endTime[i-1]
	}
	// 最大的三个间隔的下标,求的是左边的间隔
	a, b, c := 0, -1, -1
	for i := 1; i <= n; i++ {
		sz := get(i)
		if sz > get(a) {
			a, b, c = i, a, b
		} else if b < 0 || sz > get(b) {
			b, c = i, b
		} else if c < 0 || sz > get(c) {
			c = i
		}
	}
	ans := 0

	for i := 0; i < n; i++ {
		s, e := startTime[i], endTime[i]
		sz := e - s
		if (i != a && i+1 != a && get(a) >= sz) ||
			(i != b && i+1 != b && get(b) >= sz) ||
			(i != c && i+1 != c && get(c) >= sz) {
			ans = max(ans, get(i)+get(i+1)+sz)
		} else {
			ans = max(ans, get(i)+get(i+1))
		}
	}
	return ans
}
