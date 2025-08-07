package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}

func dailyTemperatures2(temperatures []int) []int {
	st := make([]int, 0)
	n := len(temperatures)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && temperatures[st[len(st)-1]] <= temperatures[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}

// 8,1,5,4,3,2,1,1,0,0

// 请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：
// 要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
func dailyTemperatures(temperatures []int) []int {
	st := make([]int, 0)
	n := len(temperatures)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
