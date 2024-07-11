package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(peopleIndexes([][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}}))
}

// 1 <= favoriteCompanies[i].length <= 500 导致没有办法用二进制
func peopleIndexes(favoriteCompanies [][]string) []int {
	cnt := make(map[string]int)
	start := 1
	for _, ch := range favoriteCompanies {
		for _, j := range ch {
			if cnt[j] == 0 {
				cnt[j] = start
				start++
			}
		}
	}
	n := len(favoriteCompanies)
	fav := make([]int, n)

	for i, fa := range favoriteCompanies {
		set := 0
		for _, j := range fa {
			set = set | (1 << cnt[j])
		}
		fav[i] = set
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		yes := true
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if fav[i]&fav[j] == fav[i] {
				yes = false
				break
			}
		}
		if yes {
			ans = append(ans, i)
		}
	}
	sort.Ints(ans)
	return ans
}
