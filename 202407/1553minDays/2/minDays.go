package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDays(10))
	fmt.Println(minDays(20))
}

var mem map[int]int

func init() {
	mem = make(map[int]int)
}

// 代码正确但是会超时
func minDays1(n int) int {
	var dfs func(i int) int

	dfs = func(i int) int {
		if i <= 1 {
			return i
		}
		if va, ok := mem[i]; ok {
			return va
		}

		res := 1 + dfs(i-1)
		if i&1 == 0 {
			res = min(res, dfs(i/2)+1)
		}
		if i%3 == 0 {
			res = min(res, dfs(i-2*i/3)+1)
		}
		mem[i] = res
		return res

	}
	return dfs(n)
}

// 两种递归的差异在那里呢
func minDays(n int) int {
	memo := map[int]int{}
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return i
		}
		if v, ok := memo[i]; ok { // 之前计算过
			return v
		}
		res := min(dfs(i/2)+i%2, dfs(i/3)+i%3) + 1
		memo[i] = res // 记忆化
		return res
	}
	return dfs(n)
}
