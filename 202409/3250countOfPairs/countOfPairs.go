package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(countOfPairs([]int{2, 3, 2}))
	fmt.Println(countOfPairs([]int{5, 5, 5, 5}))
}

func countOfPairs1(nums []int) int {
	n := len(nums)
	mx := slices.Max(nums)
	mod := int(1e9 + 7)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, mx+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		cnt := 0
		maxK := min(j, nums[i-1]-nums[i]+j)
		// 要求非负数
		for k := 0; k <= maxK; k++ {
			cnt += dfs(i-1, k)
		}
		mem[i][j] = cnt % mod
		return cnt
	}
	ans := 0
	for j := nums[n-1]; j >= 0; j-- {
		ans += dfs(n-1, j)
	}
	return ans % mod
}

func countOfPairs(nums []int) int {
	n := len(nums)
	mx := slices.Max(nums)
	mod := int(1e9 + 7)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, mx+1)
	}

	for j := range f[0] {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		for j := nums[i]; j >= 0; j-- {
			maxK := min(j, nums[i-1]-nums[i]+j)
			for k := 0; k <= maxK; k++ {
				f[i][j] += f[i-1][k]
				f[i][j] = f[i][j] % mod
			}
		}
	}

	ans := 0
	for j := mx; j >= 0; j-- {
		ans += f[n-1][j]
		ans = ans % mod
	}

	return ans % mod
}
