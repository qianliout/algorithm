package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

func combinationSum42(nums []int, target int) int {
	var dfs func(x int) int
	mem := make([]int, target+1)
	for i := range mem {
		mem[i] = -1
	}

	dfs = func(sum int) int {
		if sum == 0 {
			return 1
		}
		if sum < 0 {
			return 0
		}
		if mem[sum] != -1 {
			return mem[sum]
		}
		ans := 0
		for _, ch := range nums {
			ans += dfs(sum - ch)
		}
		mem[sum] = ans
		return ans
	}
	ans := dfs(target)
	return ans
}

func combinationSum4(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				f[i] += f[i-num]
			}
		}
	}
	return f[target]
}
