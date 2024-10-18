package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(knightDialer(3131))
}

func knightDialer(n int) int {
	MOD := int(math.Pow10(9)) + 7
	moves := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {3, 9, 0}, {}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4}}

	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 10)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(i int, num int) int

	dfs = func(i int, num int) int {
		if i == n-1 {
			return 1
		}
		if mem[i][num] != -1 {
			return mem[i][num]
		}
		res := 0
		for _, ch := range moves[num] {
			res += dfs(i+1, ch)
		}
		mem[i][num] = res % MOD
		return res % MOD
	}
	res := 0
	for i := 0; i < 10; i++ {
		res += dfs(0, i)
	}
	return res % MOD
}
