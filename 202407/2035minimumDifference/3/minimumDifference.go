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
	// fmt.Println(minimumDifference([]int{76, 8, 45, 20, 74, 84, 28, 1}))
	fmt.Println(minimumDifference([]int{7772197, 4460211, -7641449, -8856364, 546755, -3673029, 527497, -9392076, 3130315, -5309187, -4781283, 5919119, 3093450, 1132720, 6380128, -3954678, -1651499, -7944388, -3056827, 1610628, 7711173, 6595873, 302974, 7656726, -2572679, 0, 2121026, -5743797, -8897395, -9699694}))
}

// 这样写又会超时
func minimumDifference(nums []int) int {
	n := len(nums)
	m := 1 << n
	mask := make([]int, 0)
	for i := 0; i < m; i++ {
		if bits.OnesCount(uint(i)) == n/2 {
			mask = append(mask, cal(i, nums))
		}
	}
	all := 0
	for _, ch := range nums {
		all += ch
	}
	ans := math.MaxInt64 / 10
	// 你需要将 nums 分成 两个 长度为 n/2 的数组
	for _, ch := range mask {
		ans = min(ans, abs(ch-(all-ch)))
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
