package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, c := range nums {
		sum += c
	}
	if sum&1 == 1 {
		return false
	}
	sum = sum / 2
	n := len(nums)
	mem := make([]map[int]bool, n+1)

	ans := dfs(nums, 0, sum, mem)
	return ans
}

/*
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

// 面试时最好直接写 dfs 容易想，也不容易错
func dfs(nums []int, start int, target int, mem []map[int]bool) bool {
	//  这一步一定要在最前面判断，因为下面没有判断 start+1这些操作
	//  这样判断边界条件不好
	if target == 0 {
		return true
	}
	if start >= len(nums) {
		return false
	}
	if target < 0 {
		return false
	}

	if va, ok := mem[start][target]; ok {
		return va
	}

	yes := dfs(nums, start+1, target-nums[start], mem)
	no := dfs(nums, start+1, target, mem)
	if mem[start] == nil {
		mem[start] = make(map[int]bool)
	}
	mem[start][target] = yes || no
	return yes || no
}

func dfs2(nums []int, i int, target int, mem []map[int]bool) bool {
	if i >= len(nums) {
		return target == 0
	}

	if va, ok := mem[i][target]; ok {
		return va
	}

	yes := dfs2(nums, i+1, target-nums[i], mem)
	no := dfs2(nums, i+1, target, mem)
	if mem[i] == nil {
		mem[i] = make(map[int]bool)
	}
	mem[i][target] = yes || no
	return yes || no
}

func dfs3(nums []int, i int, target int, mem []map[int]bool) bool {
	if i < 0 {
		return target == 0
	}

	if va, ok := mem[i][target]; ok {
		return va
	}

	yes := dfs3(nums, i-1, target-nums[i], mem)
	no := dfs3(nums, i-1, target, mem)
	if mem[i] == nil {
		mem[i] = make(map[int]bool)
	}
	mem[i][target] = yes || no
	return yes || no
}

func canPartition(nums []int) bool {
	sum := 0
	for _, c := range nums {
		sum += c
	}
	if sum&1 == 1 {
		return false
	}
	sum = sum / 2
	n := len(nums)
	f := make([][]bool, n+5)
	for i := range f {
		f[i] = make([]bool, sum+10)
	}

	// f[i][c] = f[i-1][c] || f[i-1][c-nums[i]]
	//  防止下标越界
	// f[i+1][c] = f[i][c] || f[i][c-nums[i]]

	// 初值
	f[0][0] = true
	// 面试时两树组的方式能写对就行了
	for i := 0; i < n; i++ {
		for c := 0; c <= sum; c++ {
			f[i+1][c] = f[i][c] || (c >= nums[i] && (f[i][c-nums[i]]))
		}
	}

	return f[n][sum]
}
