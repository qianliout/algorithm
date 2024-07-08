package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
	fmt.Println(rankTeams([]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}))
}

func rankTeams(votes []string) string {
	pairs := make(map[string]*pair)
	for _, ch := range votes {
		for i, by := range ch {
			if pairs[string(by)] == nil {
				pairs[string(by)] = &pair{Apl: string(ch), Data: make([]int, 26)}
			}
			c := pairs[string(by)]
			c.Data[i]++
			pairs[string(by)] = c
		}
	}
	ans := make([]string, 0)
	for k := range pairs {
		ans = append(ans, k)
	}
	// 参赛团队的排名次序依照其所获「排位第一」的票的多少决定。如果存在多个团队并列的情况，将继续考虑其「排位第二」的票的数量。以此类推，直到不再存在并列的情况。
	sort.Slice(ans, func(i, j int) bool {
		a := pairs[ans[i]].Data
		b := pairs[ans[j]].Data
		for k := 0; k < 26; k++ {
			if a[k] > b[k] {
				return true
			} else if a[k] < b[k] {
				return false
			}
		}
		return ans[i] < ans[j]
	})
	return strings.Join(ans, "")
}

type pair struct {
	Apl  string
	Data []int
}
