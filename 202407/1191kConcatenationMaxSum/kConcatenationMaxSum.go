package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(kConcatenationMaxSum([]int{1, 2}, 3))
	fmt.Println(kConcatenationMaxSum([]int{10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000}, 100000))
}

var mod = int(math.Pow10(9)) + 7

func kConcatenationMaxSum(arr []int, k int) int {
	if k <= 0 {
		return 0
	}
	if k == 1 {
		return help(arr)
	}
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	arr = append(arr, arr...)
	if sum > 0 {
		return (help(arr) + (k-2)*sum) % mod
	}
	return help(arr) % mod
}

// 当 k==1的时候
func help(nums []int) int {
	ans := 0
	dp := make([]int, len(nums))
	for i, ch := range nums {
		dp[i] = max(dp[i], 0, ch)
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1]+ch)
		}
		ans = max(ans, dp[i])
	}
	return ans % mod
}
