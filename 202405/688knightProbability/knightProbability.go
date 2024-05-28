package main

func main() {

}

// 这里的方向一定要写对
var dirs = [][]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, 2}, {1, 2}, {-1, -2}, {1, -2}}

func knightProbability(n int, k int, row int, column int) float64 {
	// 定义 f[i][j][p] 为从位置 (i,j) 出发，使用步数不超过 p 步，最后仍在棋盘内的概率。
	dp := make([][][]float64, n)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, k+1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][0] = 1 // 移动0次，概率是1
		}
	}

	for p := 1; p <= k; p++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for _, dir := range dirs {
					nx, ny := i+dir[0], j+dir[1]
					if nx < 0 || nx >= n || ny < 0 || ny >= n {
						continue
					}
					dp[i][j][p] += dp[nx][ny][p-1] / 8
				}
			}
		}
	}

	return dp[column][row][k]
}
