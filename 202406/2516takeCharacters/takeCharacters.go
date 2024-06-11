package main

import (
	"fmt"
)

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
}

func takeCharacters(s string, k int) int {
	cnt := make([]int, 3)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, v := range cnt {
		if v < k {
			return -1
		}
		cnt[i] = v - k
	}

	le, ri, ans, n := 0, 0, -1, len(s)

	need := make([]int, 3)

	for le <= ri && ri < n {
		c := s[ri] - 'a'
		need[c]++
		ri++
		for !check(need, cnt) {
			need[s[le]-'a']--
			le++
		}
		if check(need, cnt) {
			ans = max(ans, ri-le)
		}
	}
	if ans == -1 {
		return ans
	}
	return n - ans
}

func check(cnt []int, need []int) bool {
	for i, c := range cnt {
		if c > need[i] {
			return false
		}
	}
	return true
}
