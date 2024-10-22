package main

func main() {

}

func new21Game(n int, k int, maxPts int) float64 {
	// dp[x] 表示他手上的牌总数是 x 时获胜的概率（当停止抽牌时，手上总数大于 n 就是获胜）
	dp := make([]float64, k+maxPts)
	// dp[x]=1/w * (dp[x+1]+dp[x+2]+dp[x+3]...+dp[x+w])
	s := float64(0)
	// 当x>=K时，爱丽丝会停止抽牌，这个时候游戏已经结束了，她是赢是输也已经确定了，所以此时赢的概率要么1，要么0
	for i := k + maxPts; i >= k; i-- {
		// 结束时她所持牌面值小于等于N的概率
		if i <= n {
			dp[i] = 1
		}
		s += dp[i]
	}
	for i := k - 1; i >= 0; i-- {
		dp[i] = s / float64(maxPts)
		// 这一段还没有能理解
		s = s - dp[i+maxPts] + dp[i]
	}
	return dp[0]
}
