package main

import (
	"fmt"
)

func main() {
	// fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}

func minMutation1(startGene string, endGene string, bank []string) int {
	bank2 := make(map[string]bool)
	for _, ch := range bank {
		bank2[ch] = true
	}
	if !bank2[endGene] {
		return -1
	}
	q := []string{startGene}
	visit := make(map[string]bool)
	visit[startGene] = true
	ans := 0
	for len(q) > 0 {
		lve := make([]string, 0)
		for _, no := range q {
			if no == endGene {
				return ans
			}
			nex := help(no, bank2)
			for _, ch := range nex {
				if !visit[ch] {
					lve = append(lve, ch)
					visit[ch] = true
				}
			}
		}
		if len(lve) > 0 {
			ans++
		}
		q = lve

	}
	return -1
}

func help(s string, cnt map[string]bool) []string {
	ans := make([]string, 0)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		pre := ss[i]
		for _, ch := range []byte{'A', 'C', 'G', 'T'} {
			if pre == ch {
				continue
			}
			ss[i] = ch
			if cnt[string(ss)] {
				ans = append(ans, string(ss))
			}
		}
		ss[i] = pre
	}
	return ans
}

func minMutation(startGene string, endGene string, bank []string) int {
	bank2 := make(map[string]bool)
	for _, ch := range bank {
		bank2[ch] = true
	}
	if !bank2[endGene] {
		return -1
	}
	q := []string{startGene}
	ans := 0
	for len(q) > 0 {
		lve := make([]string, 0)
		for _, no := range q {
			if no == endGene {
				return ans
			}

			nex := help(no, bank2)
			for _, ch := range nex {
				lve = append(lve, ch)
				bank2[ch] = false
			}
		}
		if len(lve) > 0 {
			ans++
		}
		q = lve
	}
	return -1
}
