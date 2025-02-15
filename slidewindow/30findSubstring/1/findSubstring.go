package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	// fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	fmt.Println(findSubstring2("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	// fmt.Println(findSubstring("fooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}

func findSubstring2(s string, words []string) []int {
	ans := make([]int, 0)
	if len(words) == 0 || len(words[0]) == 0 {
		return ans
	}
	m, n := len(words[0]), len(s)
	aa := make(map[string]int)
	for _, ch := range words {
		aa[ch]++
	}
	ll := m * len(words)
	for i := 0; i < m; i++ {
		win := make(map[string]int)
		le := i
		ri := i
		for le <= ri && ri+m <= n {
			w := s[ri : ri+m]
			win[w]++
			ri += m

			if le+ll == ri && check(aa, win) {
				ans = append(ans, le)
			}
			// 出窗口的条件是 le+ll<=ri 这个判断很重要
			for le+m <= n && le+ll <= ri {
				left := s[le : le+m]
				win[left]--
				le += m
			}
		}
	}
	// 去重
	res := make([]int, 0)
	used := make(map[int]bool)
	for _, ch := range ans {
		if !used[ch] {
			used[ch] = true
			res = append(res, ch)
		}
	}
	return ans
}

func check(aa map[string]int, bb map[string]int) bool {
	for k, v := range aa {
		if bb[k] != v {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	if len(words) == 0 || len(words[0]) == 0 {
		return ans
	}
	m, n := len(words[0]), len(s)
	wordCount := len(words)
	// totalLen := m * wordCount
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < m; i++ {
		left := i
		right := i
		currentMap := make(map[string]int)
		count := 0

		for right+m <= n {
			word := s[right : right+m]
			right += m

			if _, ok := wordMap[word]; ok {
				currentMap[word]++
				count++

				for currentMap[word] > wordMap[word] {
					leftWord := s[left : left+m]
					currentMap[leftWord]--
					count--
					left += m
				}

				if count == wordCount {
					ans = append(ans, left)
				}
			} else {
				currentMap = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return ans
}
