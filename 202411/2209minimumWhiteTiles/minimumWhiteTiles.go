package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumWhiteTiles("10110101", 2, 2))
	fmt.Println(minimumWhiteTiles("11111", 2, 3))
}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	ss := []byte(floor)
	n := len(ss)
	sum := make([]int, n+1)
	for i, ch := range ss {
		sum[i+1] = sum[i] + int(ch-'0')
	}
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, numCarpets+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		if j <= 0 {
			return sum[i+1]
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := 0
		if ss[i] == '0' {
			res = dfs(i-1, j)
		} else {
			res = min(dfs(i-1, j)+1, dfs(i-carpetLen, j-1))
		}
		mem[i][j] = res
		return res
	}

	return dfs(n-1, numCarpets)
}
