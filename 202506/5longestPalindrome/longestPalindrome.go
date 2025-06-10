package main

func main() {

}

func longestPalindrome(s string) string {
	n := len(s)
	// dp := make([][]bool, n)
	// for i := range dp {
	// 	dp[i] = make([]bool, n)
	// }
	// for i := 0; i < n; i++ {
	// 	dp[i][i] = true
	// }
	ss := []byte(s)
	ans := ""
	for i := 0; i < len(s); i++ {
		a := help(ss, i, i) // 奇数长度
		if len(a) > len(ans) {
			ans = a
		}
		if i+1 < n && ss[i] == ss[i+1] {
			b := help(ss, i, i+1) // 偶数长度
			if len(b) > len(ans) {
				ans = b
			}
		}
	}

	return ans
}

func help(ss []byte, i, j int) string {
	le, ri := i, j
	ans := string(ss[i : j+1])
	for le >= 0 && ri < len(ss) {
		if ss[le] == ss[ri] {
			ans = string(ss[le : ri+1])
			le--
			ri++
		} else {
			break
		}
	}
	return ans
}
