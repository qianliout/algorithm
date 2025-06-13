package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4([]int{2, 1, 3}, 35))
}

// 一定要加 cache
func combinationSum4DFS(nums []int, target int) int {
	n := len(nums)
	var dfs func(t int) int
	mem := make([]int, target+10)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(t int) int {
		if t == 0 {
			return 1
		}
		// 1 <= nums[i] <= 1000
		if t < 0 {
			return 0 // 没有负数，可以这么判断
		}
		if mem[t] != -1 {
			return mem[t]
		}
		ans := 0
		for i := 0; i < n; i++ {
			ans += dfs(t - nums[i])
		}
		mem[t] = ans
		return ans
	}
	ans := dfs(target)
	return ans
}

/*
本质上是 70. 爬楼梯，每次从 nums 中选一个数，作为往上爬的台阶数，计算爬 target 个台阶有多少种方案。70 那题相当于 nums=[1,2]，因为每次只能爬 1 个或 2 个台阶。
*/
func combinationSum4(nums []int, target int) int {
	n := len(nums)
	// 理解f 的本质，f[i]是指要凑成 i 这个数据有多少种方法，f[0]只能是一种方法，也就是空集
	f := make([]int, target+1)
	f[0] = 1
	for t := 1; t <= target; t++ {
		for i := 0; i < n; i++ {
			if t-nums[i] >= 0 {
				f[t] += f[t-nums[i]]
			}
		}
	}
	return f[target]
}

/*
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。
*/
