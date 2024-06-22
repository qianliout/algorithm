package main

import (
	"math"
	"sort"
)

func main() {

}

func sumSubseqWidths(nums []int) int {
	mod := int(math.Pow10(9)) + 7
	n := len(nums)
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
	ans := 0
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		// nums[i]做为最大值的子序列个数是 2^i
		// nums[i]做为最小值的子序列的个数是2^(n-i-i)
		// 那这里为啥是相减呢？
		ans += (pow2[i] - pow2[n-i-1]) * nums[i] % mod
	}
	// 上面有减法，ans 可能是负数，所以不能直接返回ans
	return (ans%mod + mod) % mod
}
