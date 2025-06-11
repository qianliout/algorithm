package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 5))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1, 0}, 1)) // 2
}

/*
给你一个非负整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
*/

func findTargetSumWays2(nums []int, target int) int {
	for _, c := range nums {
		target += c
	}
	/*
		a+b=sum
		a-b=target
		a = (sum+target)/2
	*/
	if target&1 == 1 {
		return 0
	}
	target = target / 2
	mem := make([]map[int]int, len(nums)+10)
	// n := len(nums)
	ans := dfs2(nums, 0, target, mem)
	return ans
}

/*
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000

返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
*/

// 也是错的
func dfs2(nums []int, i int, target int, mem []map[int]int) int {
	if i >= len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}

	if mem[i] != nil {
		if va, ok := mem[i][target]; ok {
			return va
		}
	}
	// 不选
	ans := dfs2(nums, i+1, target, mem)
	if nums[i] <= target {
		// 按道理 target 是可以是负数的,比如-1-1+4也能得到结果2，在前面那个数时就是负数
		// 但是本题是找所有前面加加号的数的和，所以不可能是负数
		//  选
		ans += dfs2(nums, i+1, target-nums[i], mem)
	}

	if mem[i] == nil {
		mem[i] = make(map[int]int)
	}
	mem[i][target] = ans
	return ans
}

func dfs3(nums []int, i int, target int, mem []map[int]int) int {
	if i >= len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}

	if mem[i] != nil {
		if va, ok := mem[i][target]; ok {
			return va
		}
	}
	if mem[i] == nil {
		mem[i] = make(map[int]int)
	}
	// 	可以不加这个判断,因为上面只有当 target==0时才会加值，不加这个判断，会有许多无用的计算
	if nums[i] > target {
		ans := dfs3(nums, i+1, target, mem)
		mem[i][target] = ans
		return ans
	}

	ans := dfs3(nums, i+1, target, mem) + dfs3(nums, i+1, target-nums[i], mem)

	mem[i][target] = ans
	return ans
}

func findTargetSumWays(nums []int, target int) int {
	for _, c := range nums {
		target += c
	}
	if target&1 == 1 {
		return 0
	}
	target = target / 2
	//  这个判断容易
	if target < 0 {
		return 0
	}
	n := len(nums)
	f := make([][]int, n+5)
	for i := range f {
		f[i] = make([]int, target+5)
	}
	// f[i][c] = f[i-1][c]+f[i-1][c-nums[i]]
	// 为了避免下标越界
	// f[i+1][c] = f[i][c]+f[i][c-nums[i]]
	f[0][0] = 1

	for i := 0; i < n; i++ {
		for c := 0; c <= target; c++ {
			f[i+1][c] = f[i][c]
			if c-nums[i] >= 0 {
				f[i+1][c] += f[i][c-nums[i]]
			}
		}
	}

	return f[n][target]
}
