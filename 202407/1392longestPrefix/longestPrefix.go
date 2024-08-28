package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPrefix("level"))
	fmt.Println(longestPrefix("ababab"))
}

func longestPrefix(text string) string {
	hash, p := preHash(text)
	ans := ""
	n := len(text)
	mod := int64(math.Pow10(9)) + 7
	for i := 0; i <= n; i++ {
		if i == 0 || n-i <= 0 {
			continue
		}

		left := hash[i]
		right := (hash[n] - hash[n-i]*p[i]%mod + mod) % mod
		if left == right && len(ans) < i {
			ans = text[:i]
		}
	}

	return ans
}

func reverse(text string) string {
	ss := []byte(text)
	l, r := 0, len(text)-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}
	return string(ss)
}

func preHash(text string) ([]int64, []int64) {
	n := len(text)
	hash := make([]int64, n+1)
	p := make([]int64, n+1)
	p[0] = 1
	var base int64 = 27
	mod := int64(math.Pow10(9)) + 7
	for i, ch := range text {
		hash[i+1] = (hash[i]*base + encode(byte(ch))) % mod
		p[i+1] = p[i] * base % mod
	}
	return hash, p
}

func encode(ch byte) int64 {
	return int64(ch) - int64('a') + 1
}
