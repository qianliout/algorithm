package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	fmt.Println(string([]byte("lingmindraboofooowingdingbarrwingmonkeypoundcake")[13:]))
}
func findSubstring(s string, words []string) []int {
	exit := make(map[string]int)
	for _, ch := range words {
		exit[ch]++
	}
	m, n := len(words), len(words[0])

	ans := make([]int, 0)
	for i := 0; i+m*n <= len(s); i++ {
		if check(s[i:i+m*n], n, exit) {
			ans = append(ans, i)
		}
	}
	return ans
}

func check(a string, m int, b map[string]int) bool {
	cnt := make(map[string]int)
	for i := 0; i+m <= len(a); i = i + m {
		cnt[a[i:i+m]]++
	}

	for k, v := range cnt {
		if v == 0 {
			continue
		}
		if b[k] != v {
			return false
		}
	}
	for k, v := range b {
		if v == 0 {
			continue
		}
		if cnt[k] != v {
			return false
		}
	}
	return true
}
