package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	tt := make([]int, 26)
	for _, c := range p {
		idx := int(c) - int('a')
		tt[idx]++
	}
	wind := make([]int, 26)
	le, ri, n, m := 0, 0, len(s), len(p)
	ans := make([]int, 0)
	for le <= ri && ri < n {
		idx := int(s[ri]) - int('a')
		wind[idx]++
		ri++
		for le <= ri && check(tt, wind) {
			//  子串
			if ri-le == m {
				ans = append(ans, le)
			}
			c := int(s[le]) - int('a')
			wind[c]--
			le++
		}
	}
	return ans
}

func check(tt, wind []int) bool {
	for i, c := range tt {
		if wind[i] < c {
			return false
		}
	}
	return true
}
