package main

import (
	"math"
)

func main() {

}

func minimumIncrements(nums []int, target []int) int {
	m := len(target)
	inf := math.MaxInt / 10
	lms := make([]int, 1<<m)
	for i := range lms {
		lms[i] = 1
	}
	for i, ch := range target {
		b := 1 << i
		for mak := 0; mak < b; mak++ {
			lms[b|mak] = lcm(ch, lms[mak])
		}
	}
	var dfs func(i, j int) int
	mem := make([][]int, len(nums)+1)
	for i := range mem {
		mem[i] = make([]int, 1<<m+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if j == 0 {
			// 空集合了，那么就不需要做啥了
			return 0
		}
		if i < 0 {
			return inf
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		// 不修改 nums[i]
		res := dfs(i-1, j)
		// 修改nums[i]
		sub := j
		for sub > 0 {
			l := lms[sub]
			res = min(res, dfs(i-1, j^sub)+(l-nums[i]%l)%l)
			sub = (sub - 1) & j
		}
		mem[i][j] = res
		return res
	}
	res := dfs(len(nums)-1, (1<<m)-1)
	return res
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
