package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumTotalCost([]int{1, -2, 3, 4}))
}

func maximumTotalCost1(nums []int) int64 {
	var dfs func(i, j int) int
	n := len(nums)
	// 是指上一数执行后，i如果拼接到上一组后， j==0 不变号，j==1变号
	inf := math.MinInt64 / 2
	mem1 := make([][]int, n)
	for i := range mem1 {
		mem1[i] = make([]int, 2)
		for j := range mem1[i] {
			mem1[i][j] = inf
		}
	}

	dfs = func(i, j int) int {
		// i 走到了最后
		if i == n {
			return 0
		}
		if mem1[i][j] != inf {
			return mem1[i][j]
		}
		if j == 0 {
			// i不变号，那么
			// 1：i 开始自成一组，那么 i+1就要开始变号
			ans1 := dfs(i+1, 1) + nums[i]
			// 2：i加到之前的分组里，i 不变号那么 i+1就得变号
			// ans2:= dfs(i+1, 1)+nums[i]
			mem1[i][j] = ans1
			return ans1
		}

		if j == 1 {
			// 从上一步执行后，如果还是拼接到上一组中，i 需要变号
			// 那么，i 可能是第一个数，也就是i自成一组，
			ans1 := dfs(i+1, 1) + nums[i]
			// 也可以拼接到上一组中
			ans2 := dfs(i+1, 0) - nums[i]
			mem1[i][j] = max(ans1, ans2)
			return max(ans2, ans1)
		}
		return 0
	}
	ans := dfs(0, 0)
	return int64(ans)
}

func maximumTotalCost(nums []int) int64 {
	var dfs func(i int) int
	n := len(nums)
	inf := math.MinInt64 / 2
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = inf
	}
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return nums[0]
		}

		if mem[i] != inf {
			return mem[i]
		}
		/*
			对于长度超过 2 的子数组，可以继续分割为长度为 2 和 1 的子数组，而不改变成本之和。
			分成长为 1 的子数组，即 a[i] 单独作为一个长为 1 的子数组，接下来需要解决的问题为：a[0] 到 a[i−1] 的最大成本和，即 dfs(i)=dfs(i−1)+a[i]。
			分成长为 2 的子数组，即 a[i−1] 和 a[i] 作为一个长为 2 的子数组，接下来需要解决的问题为：a[0] 到 a[i−2] 的最大成本和，即 dfs(i)=dfs(i−2)+a[i−1]−a[i]。
		*/
		ans := max(dfs(i-1)+nums[i], dfs(i-2)+nums[i-1]-nums[i])
		mem[i] = ans
		return ans
	}

	return int64(dfs(n - 1))
}

func maximumTotalCost2(nums []int) int64 {
	n := len(nums)
	f := make([]int, n+1)
	f[1] = nums[0]
	for i := 1; i < n; i++ {
		f[i+1] = max(f[i]+nums[i], f[i-1]+nums[i-1]-nums[i])
	}

	return int64(f[n])
}
