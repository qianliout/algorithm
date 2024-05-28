package main

import (
	"fmt"
	"math"

	. "outback/algorithm/common/utils"
)

func main() {
	fmt.Println(findPermutation([]int{1, 0, 2}))
	fmt.Println(findPermutation([]int{0, 1, 2}))
	fmt.Println(findPermutation([]int{0, 2, 1}))
	fmt.Println(findPermutation([]int{1, 4, 8, 7, 6, 10, 3, 5, 11, 9, 0, 2}))
}

func findPermutation(nums []int) []int {
	ans := make([]int, 0)
	MakeAns(nums, 0|1<<0, 0, &ans)
	return ans
}

func MakeAns(nums []int, s int, j int, ans *[]int) {
	*ans = append(*ans, j)
	if s == (1<<len(nums))-1 {
		return
	}
	mem := make(map[int][]int)
	fi := dfs(nums, s, j, mem)

	for k := 1; k < len(nums); k++ {
		if (1<<k)&s == 0 && dfs(nums, s|(1<<k), k, mem)+AbsSub(j, nums[k]) == fi {
			MakeAns(nums, s|(1<<k), k, ans)
			break
		}
	}
}

func dfs(nums []int, s int, j int, mem map[int][]int) int {
	if len(mem[s]) > j && mem[s][j] != -1 {
		return mem[s][j]
	}

	if s == (1<<len(nums))-1 {
		if mem[s] == nil {
			mem[s] = make([]int, len(nums))
			for k := range mem[s] {
				mem[s][k] = -1
			}
		}
		ans := AbsSub(j, nums[0])
		mem[s][j] = ans
		return ans
	}

	res := math.MaxInt32
	for k := 1; k < len(nums); k++ {
		if (1<<k)&s == 0 {
			res = min(res, dfs(nums, s|(1<<k), k, mem)+AbsSub(j, nums[k]))
		}
	}

	if len(mem[s]) == 0 {
		mem[s] = make([]int, len(nums))
		for k := range mem[s] {
			mem[s][k] = -1
		}
	}
	mem[s][j] = res

	return res
}
