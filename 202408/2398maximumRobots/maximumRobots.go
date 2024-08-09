package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumRobots([]int{19, 63, 21, 8, 5, 46, 56, 45, 54, 30, 92, 63, 31, 71, 87, 94, 67, 8, 19, 89, 79, 25},
		[]int{91, 92, 39, 89, 62, 81, 33, 99, 28, 99, 86, 19, 5, 6, 19, 94, 65, 86, 17, 10, 8, 42}, 85))
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	st := make([]int, 0)
	le, ri := 0, 0
	n := len(chargeTimes)
	ans := 0
	sum := 0
	for le <= ri && ri < n {
		// 单调栈,栈顶是最大值
		for len(st) > 0 && chargeTimes[st[len(st)-1]] < chargeTimes[ri] {
			st = st[:len(st)-1]
		}
		st = append(st, ri)
		sum += runningCosts[ri]
		for le <= ri && len(st) > 0 && int64(chargeTimes[st[0]]+(ri-le+1)*sum) > budget {
			// 没有这个判断还不行，为啥呢
			// 为啥会有这个判断呢
			if st[0] == le {
				st = st[1:]
			}
			sum -= runningCosts[le]
			le++
		}
		ans = max(ans, ri-le+1)

		ri++
	}
	return ans
}
