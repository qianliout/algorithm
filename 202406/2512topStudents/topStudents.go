package main

import (
	"sort"
	"strings"
)

func main() {

}

func topStudents(pos []string, neg []string, report []string, stu []int, k int) []int {
	posM := make(map[string]int)
	negM := make(map[string]int)
	for _, ch := range pos {
		posM[ch] = 3
	}
	for _, ch := range neg {
		negM[ch] = 1
	}
	piers := make([]pier, 0)
	for i, ch := range report {
		cnt := 0
		word := strings.Split(ch, " ")
		for _, wo := range word {
			cnt += posM[wo]
			cnt -= negM[wo]
		}
		piers = append(piers, pier{id: stu[i], cnt: cnt})
	}

	sort.Slice(piers, func(i, j int) bool {
		if piers[i].cnt > piers[j].cnt {
			return true
		} else if piers[i].cnt < piers[j].cnt {
			return false
		}
		return piers[i].id < piers[j].id
	})
	ans := make([]int, 0)
	for i := 0; i < k && i < len(piers); i++ {
		ans = append(ans, piers[i].id)
	}
	return ans
}

type pier struct {
	id, cnt int
}
