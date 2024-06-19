package main

import (
	"fmt"
)

func main() {
	// fmt.Println(splitArraySameAverage([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	// fmt.Println(splitArraySameAverage([]int{0}))
	// fmt.Println(splitArraySameAverage([]int{0, 0, 3, 9, 8}))
	fmt.Println(splitArraySameAverage([]int{0, 0, 0, 0, 0}))
	// fmt.Println(splitArraySameAverage([]int{60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}))
}

/*
设数组的平均数为avg，则该问题相当于在数组中取k个数，使得其和为k * avg。对应0-1背包问题为背包容量为k*avg的情况下，能不能取k件物品把它装满。
背包总容量可以定为数组的和
*/
func splitArraySameAverage1(nums []int) bool {
	sum, n := 0, len(nums)
	for _, ch := range nums {
		sum += ch
	}
	if sum == 0 {
		return true
	}
	// 用32位二进制的每一位来表示可选的k值，即第k位为1就表示可以选k个数字。计数+1操作就可以通过左移一位来实现。
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := sum / 2; j > 0 && j > nums[i]; j-- {
			dp[j] = dp[j] | dp[j-nums[i]]<<1
			if (j*n)%sum == 0 && (1<<(j*n%sum))&dp[j] > 0 {
				return true
			}
		}
	}
	return false
}

// 01
func splitArraySameAverage(nums []int) bool {

	sum, n := 0, len(nums)
	for _, ch := range nums {
		sum += ch
	}
	// [0,0,0] 这种情况下，面面j的逻辑都不会走
	if sum == 0 && n >= 2 {
		return true
	}

	dp := make([]map[int]bool, sum+1)
	dp[0] = map[int]bool{}
	for i := 0; i < n; i++ {
		for j := sum - 1; j >= nums[i]; j-- {
			if dp[j-nums[i]] == nil {
				continue
			}
			if dp[j] == nil {
				dp[j] = make(map[int]bool)
			}
			if j == nums[i] {
				dp[j][1] = true
			}
			pre := dp[j-nums[i]]
			for c := range pre {
				dp[j][c+1] = true
			}
			// 判断J 了
			for c := range dp[j] {
				if j*n == c*sum {
					return true
				}
			}
		}
	}
	return false
}
