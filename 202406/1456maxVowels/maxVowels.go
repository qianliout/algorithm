package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func maxVowels(s string, k int) int {
	ans, cnt := 0, 0
	le, ri, n := 0, 0, len(s)
	for le <= ri && ri < n {
		cnt += check(s[ri])
		ri++
		for ri-le >= k {
			ans = max(ans, cnt)
			cnt -= check(s[le])
			le++
		}
	}
	return ans
}

func check(b byte) int {
	mm := make(map[byte]bool)
	mm['a'] = true
	mm['e'] = true
	mm['i'] = true
	mm['o'] = true
	mm['u'] = true
	if mm[b] {
		return 1
	}
	return 0
}
