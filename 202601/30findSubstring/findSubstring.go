package main

import (
	"fmt"
)

func main() {
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo", "barr", "wing", "ding", "wing"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring2(s string, words []string) []int {
	n, m := len(s), len(words[0])
	all := m * len(words)

	tt := make(map[string]int)
	for _, w := range words {
		tt[w]++
	}

	ans := make([]int, 0)
	// for i:=0;i<n;i++{
	// 这里i的范围是关建
	for i := 0; i < m; i++ {
		wind := make(map[string]int)
		l, r := i, i
		for l <= r && r+m <= n {
			c := s[r : r+m]
			wind[c]++
			for l+m <= n && check(wind, tt) {
				if r+m-l == all {
					ans = append(ans, l)
				}
				wind[s[l:l+m]]--
				l = l + m
			}
			r = r + m
		}
	}
	// 去重
	ans = dep(ans)
	return ans
}

func check(a, b map[string]int) bool {
	for k, v := range b {
		if a[k] < v {
			return false
		}
	}
	return true
}

func dep(num []int) []int {
	ans := make([]int, 0)
	exit := make(map[int]int)
	for _, v := range num {
		if exit[v] == 0 {
			ans = append(ans, v)
		}
		exit[v]++
	}
	return ans
}

func findSubstring(s string, words []string) []int {
	n, m := len(s), len(words[0])
	all := m * len(words)

	tt := make(map[string]int)
	for _, w := range words {
		tt[w]++
	}

	ans := make([]int, 0)
	// for i:=0;i<n;i++{
	// 这里i的范围是关建
	for i := 0; i < m; i++ {
		wind := make(map[string]int)
		l, r := i, i
		for l <= r && r+m <= n {
			c := s[r : r+m]
			wind[c]++
			if r+m-l == all && check(wind, tt) {
				ans = append(ans, l)
			}
			for r+m-l >= all {
				cl := s[l : l+m]
				wind[cl]--
				l = l + m
			}
			r = r + m
		}
	}
	// 去重
	ans = dep(ans)
	return ans
}
