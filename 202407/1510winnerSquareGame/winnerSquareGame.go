package main

func main() {

}

func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	// 加不加初值没有影响
	dp[0] = false
	dp[1] = true
	for i := 0; i <= n; i++ {
		for j := 1; j*j+i <= n; j++ {
			if !dp[i] {
				dp[j*j+i] = true
			}
		}
	}
	return dp[n]
}
