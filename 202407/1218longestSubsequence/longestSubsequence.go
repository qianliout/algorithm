package main

func main() {
}

// 这样写会超时
func longestSubsequence1(arr []int, difference int) int {
	n := len(arr)
	dp := make([]int, n)
	ans := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if arr[i]-arr[j] == difference {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	dp := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		pre := arr[i] - difference
		dp[arr[i]] = dp[pre] + 1
		ans = max(ans, dp[arr[i]])
	}
	return ans
}
