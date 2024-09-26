package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(countOfPairs([]int{2, 3, 2}))
	fmt.Println(countOfPairs([]int{5, 5, 5, 5}))
}

func countOfPairs(nums []int) int {
	n := len(nums)
	mx := slices.Max(nums)
	mod := int(1e9 + 7)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, mx+1)
	}
	pre := make([]int, mx+1)
	for j := 0; j <= nums[0]; j++ {
		f[0][j] = 1
	}

	for i := 1; i < n; i++ {
		// 计算前缀和
		pre[0] = f[i-1][0]
		for k := 1; k <= mx; k++ {
			pre[k] = pre[k-1] + f[i-1][k]
		}

		for j := 0; j <= nums[i]; j++ {
			maxK := min(j, nums[i-1]-nums[i]+j)
			if maxK >= 0 {
				f[i][j] = pre[maxK] % mod
			}
		}
	}

	ans := 0
	for j := nums[n-1]; j >= 0; j-- {
		ans += f[n-1][j]
		ans = ans % mod
	}

	return ans % mod
}
