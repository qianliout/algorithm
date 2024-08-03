package main

func main() {

}

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	dp := make([]int64, len(prices))
	dp[0] = 1

	for i := 1; i < n; i++ {
		if prices[i]-prices[i-1] == -1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	var ans int64
	for _, ch := range dp {
		ans += ch
	}

	return ans
}
