package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func main() {
	fmt.Println(math.Pow(1.1, 15))
}

func maximumTotalDamage(power []int) int64 {
	cnt := make(map[int]int)
	for _, ch := range power {
		cnt[ch]++
	}

	power = slices.Compact(power)
	n := len(power)
	sort.Ints(power)
	var dfs func(i int) int

	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if mem[i] != -1 {
			return mem[i]
		}
		x := power[i]
		j := i
		for j > 0 && power[j-1] >= x-2 {
			j--
		}
		ans := max(dfs(i-1), dfs(j-1)+x*cnt[x])
		mem[i] = ans
		return ans
	}

	ans := dfs(n - 1)
	return int64(ans)
}
