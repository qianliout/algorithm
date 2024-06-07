package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations("1100011000", "0101001010", 2))
}

func minOperations(s1 string, s2 string, x int) int {
	if s1 == s2 {
		return 0
	}
	if len(s2) != len(s1) {
		return -1
	}
	p := make([]int, 0)
	for i := range s1 {
		if s1[i] != s2[i] {
			p = append(p, i)
		}
	}
	if len(p)%2 != 0 {
		return -1
	}
	var dfs func(i int) int

	mem := make(map[int]int)

	dfs = func(i int) int {
		if i == -1 {
			return 0
		}
		if va, ok := mem[i]; ok {
			return va
		}
		res1 := dfs(i-1) + x
		if i > 0 {
			res1 = min(res1, dfs(i-2)+(p[i]-p[i-1])*2)
		}
		mem[i] = res1
		return res1
	}
	return dfs(len(p)-1) / 2
}
