package main

import (
	"fmt"
)

func main() {
	fmt.Println(waysToPartition([]int{22, 4, -25, -20, -15, 15, -16, 7, 19, -10, 0, -13, -14}, -33))
}

func waysToPartition(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	cntR := make(map[int]int)
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i]
		// 这里为啥是 sum[i-1] 呢
		cntR[sum[i-1]]++
	}
	all := sum[n-1]
	ans := 0
	if all%2 == 0 {
		ans = cntR[all/2]
	}
	cntL := make(map[int]int)

	for i := 0; i < n; i++ {
		ch := sum[i]
		d := k - nums[i]
		if (all+d)%2 == 0 {
			ans = max(ans, cntL[(all+d)/2]+cntR[(all-d)/2])
		}
		cntL[ch]++
		cntR[ch]--
	}

	return ans
}
