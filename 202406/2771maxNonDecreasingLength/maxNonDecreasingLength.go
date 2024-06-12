package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxNonDecreasingLength([]int{2, 3, 1}, []int{1, 2, 1}))
	fmt.Println(maxNonDecreasingLength([]int{11, 7, 7, 9}, []int{19, 19, 1, 7}))
}

func maxNonDecreasingLength2(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	ans, n, cnt := 1, len(nums1), 1

	pre := min(nums1[0], nums2[0])

	for i := 1; i < n; i++ {
		if max(nums1[i], nums2[i]) < pre {
			ans = max(ans, cnt)
			cnt = 1
			pre = min(nums1[i], nums2[i])
			continue
		}
		if nums1[i] < pre {
			pre = nums2[i]
			cnt++
		} else if nums2[i] < pre {
			pre = nums1[i]
			cnt++
		} else {
			pre = min(nums1[i], nums2[i])
			cnt++
		}
	}
	ans = max(ans, cnt)
	return ans
}
func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	nums := [2][]int{nums1, nums2}
	var dfs func(i, j int) int
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 2)
	}
	dfs = func(i, j int) int {
		if i <= 0 {
			return 1
		}
		res := 1
		if mem[i][j] > 0 {
			return mem[i][j]
		}
		if nums1[i-1] <= nums[j][i] {
			res = max(res, dfs(i-1, 0)+1)
		}
		if nums2[i-1] <= nums[j][i] {
			res = max(res, dfs(i-1, 1)+1)
		}
		mem[i][j] = res
		return res
	}
	ans := 0
	for j := 0; j < 2; j++ {
		for i := 0; i < n; i++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
