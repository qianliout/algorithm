package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// f[i][j] 表示前word1前i个,word2前 j个字符（不包括 i,j）的编辑距离
	f := make([][]int, m+5)
	for i := range f {
		f[i] = make([]int, n+5)
	}
	// 初值
	// 为啥这里不写这几行的初值就会出错呢
	for i := 1; i <= m; i++ {
		f[i][0] = i
	}
	for j := 1; j <= n; j++ {
		f[0][j] = j
	}
	// 因为这里都是从1开如的，所以对于0就必须有初值
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
		}
	}
	return f[m][n]
}
