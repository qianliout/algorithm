package main

import (
	"math"
)

func main() {

}

func stoneGameIII(s []int) string {
	n := len(s)
	s[n-1] = s[n-1]
	for i := n - 2; i >= 0; i-- {
		s[i] = s[i] + s[i+1]
	}

	inf := math.MaxInt / 10
	var dfs func(i int) int
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i < 0 || i >= n {
			return 0
		}
		if mem[i] != -1 {
			return mem[i]
		}
		mi := inf
		for j := 1; j <= 3; j++ {
			mi = min(mi, dfs(i+j))
		}
		mem[i] = s[i] - mi
		return s[i] - mi
	}
	a := dfs(0)*2 - s[0]

	if a > 0 {
		return "Alice"
	} else if a < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}
