package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// fmt.Println(minimumDifference([]int{2, -1, 0, 4, -2, -9}))
	// fmt.Println(minimumDifference([]int{3, 9, 7, 3}))
	// fmt.Println(minimumDifference([]int{-36, 36}))
	fmt.Println(minimumDifference([]int{76, 8, 45, 20, 74, 84, 28, 1}))
}

// 可以得到结果，但是会超内存
func minimumDifference(nums []int) int {
	n := len(nums)
	m := 1 << n
	mask := make([]int, m)
	for i := 0; i < m; i++ {
		mask[i] = cal(i, nums)
	}
	all := 0
	for _, ch := range nums {
		all += ch
	}
	ans := math.MaxInt64 / 10
	// 你需要将 nums 分成 两个 长度为 n/2 的数组
	for i := 0; i < m; i++ {
		if bits.OnesCount(uint(i)) == n/2 {
			ans = min(ans, abs(mask[i]-(all-mask[i])))
		}
	}
	return ans
}

func cal(set int, nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if set&(1<<i) != 0 {
			ans += nums[i]
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
