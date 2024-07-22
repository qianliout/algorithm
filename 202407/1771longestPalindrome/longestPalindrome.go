package main

func main() {

}

func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	ans, n := 0, len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	// i和 j 的执行顺序是啥，怎么去理解
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1 // 初值
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
				// 需要理解的点：为啥只是需要在这里更新答案，其他地方不更新
				if i < len(word1) && j >= len(word1) {
					ans = max(ans, f[i][j])
				}
			} else {
				f[i][j] = max(f[i+1][j], f[i][j-1])
				// 这里不能更新答案，不估就会出错
				// if i < len(word1) && j >= len(word1) {
				// 	ans = max(ans, f[i][j])
				// }
			}
		}
	}
	return ans
}
