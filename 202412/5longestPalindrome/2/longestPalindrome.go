package main

func main() {

}

func longestPalindrome(s string) string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
	}
	ans := ""
	for r := 0; r < n; r++ {
		for l := r; l >= 0; l-- {
			if s[l] == s[r] && (r-l <= 2 || f[l+1][r-1]) {
				f[l][r] = true
				if len(ans) < r-l+1 {
					ans = s[l : r+1]
				}
			}
		}
	}

	return ans
}
