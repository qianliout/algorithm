package main

func main() {

}

func champagneTower(poured int, row int, glass int) float64 {
	dp := make([][]float64, row+10)
	for i := range dp {
		dp[i] = make([]float64, row+10)
	}
	dp[0][0] = float64(poured)

	for i := 0; i <= row; i++ {
		for j := 0; j <= i; j++ {
			if dp[i][j] <= 1 {
				continue
			}
			dp[i+1][j] += (dp[i][j] - 1) / 2
			dp[i+1][j+1] += (dp[i][j] - 1) / 2
		}
	}
	return min(1, dp[row][glass])
}
