package main

import (
	"fmt"
)

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}))
}

func mincostTickets1(days []int, costs []int) int {
	da := make(map[int]int)
	for _, ch := range days {
		da[ch] = 1
	}

	n := len(days)
	var dfs func(i int) int
	mx := days[n-1]
	mem := make([]int, mx+1)

	dfs = func(i int) int {
		if i <= 0 {
			return 0
		}
		if mem[i] > 0 {
			return mem[i]
		}
		ans := 0
		if _, ok := da[i]; !ok {
			ans = dfs(i - 1)
		} else {
			a := dfs(i-1) + costs[0]
			b := dfs(i-7) + costs[1]
			c := dfs(i-30) + costs[2]
			ans = min(a, b, c)
		}
		mem[i] = ans
		return ans
	}

	ans := dfs(days[n-1])
	return ans
}

// 假设第 100 天是旅行的最后一天，分类讨论：
// 在第 100 天购买为期 1 天的通行证，接下来需要解决的问题为：1 到 99 天的最小花费。
// 在第 94 天购买为期 7 天的通行证，接下来需要解决的问题为：1 到 93 天的最小花费。
// 在第 71 天购买为期 30 天的通行证，接下来需要解决的问题为：1 到 70 天的最小花费。

func mincostTickets(days []int, costs []int) int {
	da := make(map[int]int)
	for _, ch := range days {
		da[ch] = 1
	}

	n := len(days)
	mx := days[n-1]
	f := make([]int, mx+1)
	for i := range f {
		f[i] = 1 << 30
	}
	f[0] = 0
	for i := 1; i <= mx; i++ {
		if _, ok := da[i]; !ok {
			f[i] = f[i-1]
			continue
		}
		// 这里一定要注意不能这样写
		// if i >= 1 {
		// 	f[i] = min(f[i], f[i-1]+costs[0])
		// }

		f[i] = min(f[i], f[max(0, i-1)]+costs[0])
		f[i] = min(f[i], f[max(0, i-7)]+costs[1])
		f[i] = min(f[i], f[max(0, i-30)]+costs[2])
	}
	return f[mx]
}
