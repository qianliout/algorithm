package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}

// timeout
func threeSumMulti2(nums []int, target int) int {
	mod := int(math.Pow10(9)) + 6
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == target {
					ans++
				}
			}
		}
	}
	return ans % mod
}

func threeSumMulti(nums []int, target int) int {
	mod := int(math.Pow10(9)) + 7
	// 排序不排序没有影响
	// sort.Ints(nums)

	// dp[j][k]表示对于容量为j，且由k个数组成的可能数量
	dp := make([][]int, target+1)
	for i := range dp {
		dp[i] = make([]int, 4)
	}

	dp[0][0] = 1
	// 01 背包
	for _, ch := range nums {
		for j := target; j >= ch; j-- {
			for k := 3; k >= 1; k-- {
				dp[j][k] = (dp[j][k] + dp[j-ch][k-1]) % mod
			}
		}
	}
	return dp[target][3]
}
