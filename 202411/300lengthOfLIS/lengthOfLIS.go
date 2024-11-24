package main

func main() {

}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	mx := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		mx = max(mx, dp[i])
	}
	return mx
}
