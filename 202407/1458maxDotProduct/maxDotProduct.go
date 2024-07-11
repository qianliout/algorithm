package main

import (
	"math"
)

func main() {

}

func maxDotProduct(nums1 []int, nums2 []int) int {
	var dfs func(i, j int) int
	m, n := len(nums1), len(nums2)
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MinInt / 10
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := max(0, dfs(i-1, j-1)) + nums1[i]*nums2[j]
		res = max(res, dfs(i-1, j), dfs(i, j-1))
		mem[i][j] = res
		return res
	}
	res := dfs(m-1, n-1)
	return res
}
