package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1}, 1))
}

func findTargetSumWays1(nums []int, target int) int {
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	if sum+target < 0 || (sum+target)%2 == 1 {
		return 0
	}

	c := (sum + target) / 2 // y 的个数
	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, c+1)
	}
	// 初值
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= c; j++ {
			f[i][j] = f[i-1][j]
			if j >= nums[i-1] {
				f[i][j] += f[i-1][j-nums[i-1]]
			}
		}
	}
	return f[n][c]
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	if sum+target < 0 || (sum+target)%2 == 1 {
		return 0
	}

	target = (sum + target) / 2 // y 的个数

	f := make([]int, target+1)
	f[0] = 1
	// for i := 1; i <= n; i++ {
	// 	for j := target; j >= 0; j-- {
	// 		f[j] = f[j]
	// 		if j >= nums[i-1] {
	// 			f[j] += f[j-nums[i-1]]
	// 		}
	// 	}
	// }
	// 可以再精简
	for _, x := range nums {
		for c := target; c >= x; c-- {
			f[c] = f[c] + f[c-x]
		}
	}

	return f[target]
}

func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	if sum+target < 0 || (sum+target)%2 == 1 {
		return 0
	}

	target = (sum + target) / 2 // y 的个数
	n := len(nums)

	var dfs func(i, s int) int
	dfs = func(i, s int) int {
		if i < 0 {
			if s == 0 {
				return 1
			}
			return 0
		}
		ans := dfs(i-1, s)
		if s >= nums[i] {
			ans += dfs(i-1, s-nums[i])
		}
		return ans
	}
	ans := dfs(n-1, target)
	return ans
}
