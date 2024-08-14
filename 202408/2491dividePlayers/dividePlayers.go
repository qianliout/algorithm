package main

import (
	"sort"
)

func main() {

}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	n := len(skill)
	a := skill[0] + skill[n-1]
	ans := skill[0] * skill[n-1]
	for i := 1; i < n/2; i++ {
		b := skill[i] + skill[n-1-i]
		if a != b {
			return -1
		}
		ans += skill[i] * skill[n-1-i]
	}
	return int64(ans)
}
