package main

import (
	"fmt"
)

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{}))
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	bank2 := make(map[string]bool)

	used := make(map[string]bool)
	for _, ch := range bank {
		bank2[ch] = true
	}

	if !bank2[endGene] {
		return -1
	}

	bank2[endGene] = true
	used[startGene] = true
	q := []string{startGene}
	ans := 0
	for len(q) > 0 {
		lev := make([]string, 0)
		for _, no := range q {
			if no == endGene {
				return ans
			}
			for _, nx := range help(no, bank2) {
				if used[nx] {
					continue
				}

				if nx == endGene {
					return ans + 1
				}

				used[nx] = true
				lev = append(lev, nx)
			}
		}
		if len(lev) > 0 {
			ans++
		}
		q = lev

	}
	return -1
}

func help(start string, bank2 map[string]bool) []string {
	ans := make([]string, 0)
	ss := []byte(start)
	n := len(ss)
	for i := 0; i < n; i++ {
		pre := ss[i]
		for j := 'A'; j <= 'Z'; j++ {
			if pre == byte(j) {
				continue
			}
			ss[i] = byte(j)
			if !bank2[string(ss)] {
				continue
			}
			ans = append(ans, string(ss))
		}
		ss[i] = pre
	}
	return ans
}
