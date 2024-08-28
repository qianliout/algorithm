package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(distinctEchoSubstrings("leetcodeleetcode"))
	fmt.Println(distinctEchoSubstrings("leeleetleeleeleclec"))
}

func distinctEchoSubstrings(text string) int {
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
	visit := make(map[int64]bool)
	ans := 0
	for i := 0; i < n; i++ {
		for m := 1; m < n; m++ {
			if i+2*m > n {
				break
			}
			// left:= text[i:m+i]
			// right:=text[i+m,i+2m]
			left := (hash[m+i] - hash[i]*p[m]%mod + mod) % mod
			right := (hash[i+2*m] - hash[i+m]*p[m]%mod + mod) % mod
			if visit[left] {
				continue
			}
			if left == right {
				visit[left] = true
				ans++
			}
		}
	}
	return ans
}

func encode(ch byte) int64 {
	return int64(ch) - int64('a') + 1
}
