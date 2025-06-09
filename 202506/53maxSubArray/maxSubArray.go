package main

func main() {

}

func maxSubArray(nums []int) int {
	ans := nums[0]
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}
