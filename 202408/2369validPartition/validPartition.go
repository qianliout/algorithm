package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))
	fmt.Println(validPartition([]int{865579, 865579, 8935936}))
}

func validPartition(nums []int) bool {
	var dfs func(i int) bool
	n := len(nums)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) bool {
		if i < 0 {
			return false
		}
		if i >= n {
			return true
		}
		if mem[i] != -1 {
			return mem[i] == 1
		}
		if i+1 < n && nums[i] == nums[i+1] && dfs(i+2) {
			mem[i] = 1
			return true
		}
		if i+2 < n && nums[i] == nums[i+1] && nums[i+1] == nums[i+2] && dfs(i+3) {
			mem[i] = 1
			return true
		}
		if i+2 < n && nums[i]+1 == nums[i+1] && nums[i+1]+1 == nums[i+2] && dfs(i+3) {
			mem[i] = 1
			return true
		}
		mem[i] = 0
		return false
	}

	return dfs(0)
}
