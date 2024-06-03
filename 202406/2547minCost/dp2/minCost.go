package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCost([]int{1, 2, 1, 2, 1, 3, 3}, 2))
	fmt.Println(minCost([]int{1, 2, 1, 2, 1}, 2))
	fmt.Println(minCost([]int{68, 81, 76, 91, 4, 0, 50, 5, 66, 23, 27, 91, 37, 61, 2, 0, 78, 11, 76, 58, 47, 46, 91, 57, 12, 61, 12, 17, 64, 26, 92, 71, 32, 52, 6, 70, 44, 77, 50, 7, 50, 2, 77, 56, 83, 69, 41, 88, 18, 2, 80, 2, 1, 2, 1, 88, 76, 88, 3, 64, 75, 93, 5, 18, 11, 65, 75, 16, 75, 17, 91, 2, 57, 10, 3, 29, 4, 83, 75, 15, 21, 20, 93, 32, 9, 29, 71, 21, 0, 93, 66, 52, 24, 90, 53, 82, 62, 85, 52, 27, 41, 58, 30, 3, 79, 63, 41, 50, 8, 71, 61, 59, 94, 13, 32, 61, 52, 80, 55, 84, 94, 61, 4, 19, 36, 25, 77, 43, 80, 72, 38, 17, 76, 74, 2, 31, 58, 39, 55, 39, 35, 88, 54, 53, 35, 80, 45, 8, 61, 67, 3, 48, 66, 15, 9, 22, 81, 30, 15, 48, 54, 87, 59, 9, 54, 65, 74, 13, 54, 48, 78, 55, 48, 93, 88, 94, 58, 84, 36, 44, 19, 50, 71, 54, 17, 76, 27, 58, 78, 88}, 178))
}

/*
用dp[i]表示前i个元素对应的最小代价。取一个j满足0<=j<i,那么dp[i]=min(dp[i],dp[j]+trimmed(j->i)+k
根据上述状态转移方程，j应该正着取还是逆着取呢？肯定是逆着取鸭，正着取的话，trimmed(j+1 to i)这个区间的重要性无从得知，所以排除
因此推出了双重循环:i走正序，j=i-1走逆序
*/
func minCost(nums []int, k int) int {
	inf := math.MaxInt / 2
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = inf
	}
	for i := 1; i <= n; i++ {
		cnt := make(map[int]int)
		sum := 0
		for j := i - 1; j >= 0; j-- {
			cnt[nums[j]]++
			c := cnt[nums[j]]
			if c == 2 {
				sum += 2
			} else if c > 2 {
				sum += 1
			}
			dp[i] = min(dp[i], dp[j]+sum+k)
		}
	}
	return dp[n]
}

// 找出并返回拆分 nums 的所有可行方案中的最小代价。
// fixme 没有能理解
func minCost1(nums []int, k int) int {
	inf := math.MaxInt / 2
	n := len(nums)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		exit := make(map[int]int)
		uni := 0
		mn := inf
		for j := i; j >= 0; j-- {
			x := nums[j]
			if exit[x] == 0 {
				exit[x]++
				uni++
			} else if exit[x] == 1 {
				exit[x]++
				uni--
			}
			mn = min(mn, dp[j]-uni)
		}
		dp[i+1] = mn + k
	}
	return dp[n] + n
}

// 这里每次都要统计，会超时
func cal(nums []int, le, ri int, k int) int {
	exit := make(map[int]int)
	for i := le; i <= ri; i++ {
		exit[nums[i]]++
	}
	ans := k
	for _, v := range exit {
		if v >= 2 {
			ans += v
		}
	}
	return ans
}
