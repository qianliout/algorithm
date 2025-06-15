package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{10, 8}, {1, 12}, {6, 15}, {2, 18}})) // 3
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		}
		return envelopes[i][1] < envelopes[j][1]
	})
	n := len(envelopes)
	var dfs func(i int) int
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}

	dfs = func(i int) int {
		if i == 0 {
			return 1
		}

		if mem[i] != -1 {
			return mem[i]
		}

		w, h := envelopes[i][0], envelopes[i][1]
		ans := 1
		for j := i - 1; j >= 0; j-- {
			w1, h1 := envelopes[j][0], envelopes[j][1]
			if w1 < w && h1 < h {
				ans = max(ans, dfs(j)+1)
			}
		}
		mem[i] = ans
		return ans
	}
	ans := dfs(n - 1)
	return ans
}
