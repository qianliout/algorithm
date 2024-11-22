package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))
}

func waysToReachTarget(target int, types [][]int) int {
	mod := int(math.Pow10(9)) + 7
	n := len(types)
	pairs := make([]pair, n)
	for i := range types {
		pairs[i] = pair{count: types[i][0], marks: types[i][1]}
	}
	var dfs func(i int, sum int) int
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, target+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i int, sum int) int {
		if i < 0 {
			if sum == 0 {
				return 1
			}
			return 0
		}
		if sum < 0 {
			return 0
		}
		if mem[i][sum] != -1 {
			return mem[i][sum]
		}
		ans := 0
		for j := 0; j <= pairs[i].count; j++ {
			if j*pairs[i].marks > sum {
				break
			}
			ans += dfs(i-1, sum-j*pairs[i].marks)
		}
		ans = ans % mod
		mem[i][sum] = ans
		return ans
	}
	ans := dfs(n-1, target)
	return ans % mod
}

type pair struct {
	count int
	marks int
}
