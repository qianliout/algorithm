package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAnagrams("ukgqajqsuset kk hm"))
	fmt.Println(closetTarget([]string{"a", "b", "leetcode"}, "leetcode", 0))
}

const mod int = 1e9 + 7

func countAnagrams(s string) int {
	ans, mul := 1, 1
	for _, s := range strings.Split(s, " ") {
		cnt := [26]int{}
		for i, c := range s {
			cnt[c-'a']++
			mul = mul * cnt[c-'a'] % mod
			ans = ans * (i + 1) % mod
		}
	}
	return ans * pow(mul, mod-2) % mod
}

// 这样写快速幂也是可以的
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		a := pow(x, n/2)
		return a * a % mod
	}
	return (x * pow(x, n-1)) % mod
}

/*
形式上， words[i] 的下一个元素是 words[(i + 1) % n] ，而 words[i] 的前一个元素是 words[(i - 1 + n) % n] ，其中 n 是 words 的长度。
*/
func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	ans := len(words)
	for i, ch := range words {
		if ch == target {
			ans = min(ans, abs(i-startIndex), n-abs(i-startIndex))
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
