package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(minAnagramLength("abba"))
	// fmt.Println(minAnagramLength("aaaaaaaaab"))
	fmt.Println(minAnagramLength("aabbabab"))
}

func minAnagramLength(s string) int {
	n := len(s)
	for k := 1; k <= n/2+1; k++ {
		if n%k != 0 {
			continue
		}
		if check(s, k) {
			return k
		}
	}

	return n
}

func minAnagramLength2(s string) int {
	n := len(s)
	fd := find(n)
	for _, ch := range fd {
		if check(s, ch) {
			return ch
		}
	}
	return n
}

func check(s string, k int) bool {
	ss := []byte(s)
	fir := ss[:k]
	for i := 0; i+k <= len(ss); i = i + k {
		if !same(ss[i:i+k], fir) {
			return false
		}
	}
	return true
}

func same(a, b []byte) bool {

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(a) == string(b)
}

// 找一个数的因子
func find(n int) []int {
	ans := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		if n%i == 0 {
			ans = append(ans, i)
		}
		if i*i < n {
			ans = append(ans, n/i)
		}
	}
	sort.Ints(ans)
	return ans
}
