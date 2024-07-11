package main

import (
	"fmt"
)

func main() {
	fmt.Println(peopleIndexes([][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}}))
}

// 1 <= favoriteCompanies[i].length <= 500 导致没有办法用二进制
func peopleIndexes(favoriteCompanies [][]string) []int {
	n := len(favoriteCompanies)
	cnt := make([]map[string]bool, n)
	for i, ch := range favoriteCompanies {
		cnt[i] = gen(ch)
	}

	ans := make([]int, 0)
	for i := range cnt {
		add := true
		for j := range cnt {
			if i == j {
				continue
			}
			if check(cnt[i], cnt[j]) {
				add = false
				break
			}
		}
		if add {
			ans = append(ans, i)
		}
	}
	return ans
}

func gen(str []string) map[string]bool {
	cnt := make(map[string]bool)
	for _, ch := range str {
		cnt[ch] = true
	}
	return cnt
}

// 检查 a 是否是 b 的子集
func check(a, b map[string]bool) bool {
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}
