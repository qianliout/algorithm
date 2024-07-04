package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 2, 3, 4}))
}

func makeArrayIncreasing(nums1 []int, nums2 []int) int {
	sort.Ints(nums2)
	var dfs func(i, pre int) int
	inf := math.MaxInt / 10
	mem := make([]map[int]int, len(nums1))
	for i := range mem {
		mem[i] = make(map[int]int)
	}

	dfs = func(i, pre int) int {
		if i < 0 {
			return 0
		}
		if va, ok := mem[i][pre]; ok {
			return va
		}

		res := inf
		if nums1[i] < pre {
			// 可以不替换
			res = min(res, dfs(i-1, nums1[i]))
		}
		// 如果 b 中有比 pre 小的数，那么选其中最大的 b[k]，去替换 a[i]
		k := sort.SearchInts(nums2, pre) - 1
		if k >= 0 {
			res = min(res, dfs(i-1, nums2[k])+1)
		}
		mem[i][pre] = res
		return res
	}
	res := dfs(len(nums1)-1, inf)
	if res >= inf {
		return -1
	}
	return res
}
