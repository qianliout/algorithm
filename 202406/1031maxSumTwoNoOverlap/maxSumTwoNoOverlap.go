package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2))
	fmt.Println(maxSumTwoNoOverlap([]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2))
}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 左fir 右sec 的
	var help func(fir, sec int) int

	help = func(fir, sec int) int {
		ans := 0
		preMax := preSum[fir]
		for i := fir + 1; i <= n+1-sec; i++ {
			sv := preSum[i+sec-1] - preSum[i-1]
			ans = max(ans, sv+preMax)
			preMax = max(preMax, preSum[i]-preSum[i-fir])
		}
		return ans
	}

	a := help(firstLen, secondLen)
	b := help(secondLen, firstLen)
	return max(a, b)
}
