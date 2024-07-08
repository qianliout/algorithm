package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))

}

var dir = []byte{'a', 'b', 'c'}

func numberOfSubstrings(s string) int {
	le, ri, n, ans := 0, 0, len(s), 0
	cnt := make(map[byte]int)
	for le <= ri && ri < n {
		ch := byte(s[ri])
		cnt[ch]++
		ri++
		for le <= ri && check(cnt) {
			cnt[byte(s[le])]--
			le++
		}
		// 所有从0到left-1开始，到right结束的子串都是有效的
		ans += le
	}
	return ans
}

func check(cnt map[byte]int) bool {
	for _, d := range dir {
		if cnt[d] <= 0 {
			return false
		}
	}
	return true
}
