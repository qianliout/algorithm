package main

import (
	"math"
)

func main() {

}

func numFactoredBinaryTrees(arr []int) int {
	base := int(math.Pow(10, 9)) + 7
	has := make(map[int]bool)
	for _, ch := range arr {
		has[ch] = true
	}
	var dfs func(va int) int
	mem := make(map[int]int)
	dfs = func(va int) int {
		if v, ok := mem[va]; ok {
			return v
		}
		res := 1
		for _, x := range arr {
			if va%x == 0 && has[va/x] {
				res += dfs(x) * dfs(va/x)
			}
		}
		mem[va] = res
		return res
	}
	ans := 0
	for _, ch := range arr {
		ans += dfs(ch)
	}

	return ans % base
}
