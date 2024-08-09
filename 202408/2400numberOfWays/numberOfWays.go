package main

import (
	"math"
)

func main() {

}

func numberOfWays(startPos int, endPos int, k int) int {
	var dfs func(int, int) int
	mod := int(math.Pow10(9)) + 7
	mem := make(map[int]map[int]int)

	dfs = func(i int, step int) int {
		if abs(i-endPos) > step {
			return 0
		}
		if step == 0 {
			return 1
		}
		if mem[i] != nil {
			if va, ok := mem[i][step]; ok {
				return va
			}
		}
		res := dfs(i+1, step-1) + dfs(i-1, step-1)
		res = res % mod
		if mem[i] == nil {
			mem[i] = make(map[int]int)
		}
		mem[i][step] = res
		return res
	}

	return dfs(startPos, k)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
