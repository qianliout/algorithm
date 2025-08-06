package main

func main() {

}

func isMatch(s string, p string) bool {
	s, p = " "+s, " "+p
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ss, pp := s[i-1], p[j-1]
			if ss == pp || pp == '?' {
				f[i][j] = f[i-1][j-1]
			}
			if pp == '*' {
				for k := i; k >= 0; k-- {
					f[i][j] = f[i][j] || f[k][j-1]
					if f[i][j] {
						break
					}
				}
			}
		}
	}

	return f[m][n]
}

/*
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符序列（包括空字符序列）。
*/
