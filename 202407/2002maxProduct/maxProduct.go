package main

import (
	"math/bits"
)

func main() {

}

func maxProduct(s string) int {
	n := len(s)
	m := 1 << n
	f := make([]int, m)
	for i := 1; i < m; i++ {
		f[i] = cal(s, i)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			// 不相交
			if i&j == 0 {
				ans = max(ans, f[i]*f[j])
			}
		}
	}

	return ans
}

// 判断是不是回文串，如果不是返回0，如果是，返长长度
func cal(s string, j int) int {
	le, ri := 0, len(s)-1

	for le < ri {
		for le < ri && (j>>le)&1 == 0 {
			le++
		}
		for le < ri && (j>>ri)&1 == 0 {
			ri--
		}
		if s[le] != s[ri] {
			return 0
		}
		le++
		ri--
	}
	return bits.OnesCount(uint(j))
}
