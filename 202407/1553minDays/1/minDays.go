package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDays(10))
	fmt.Println(minDays(937413))
}

var mem map[int]int

func init() {
	mem = make(map[int]int)
}

// 代码正确但是会超时
func minDays(n int) int {
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
