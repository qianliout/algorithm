package main

import (
	"sort"
)

func main() {

}

func findWinners(matches [][]int) [][]int {
	win := make(map[int]int)
	lose := make(map[int]int)
	all := make(map[int]int)
	for _, ch := range matches {
		w, l := ch[0], ch[1]
		win[w]++
		lose[l]++
		all[w]++
		all[l]++
	}
	answer0 := make([]int, 0)
	answer1 := make([]int, 0)
	for k := range all {
		if lose[k] == 0 && win[k] > 0 {
			answer0 = append(answer0, k)
		}
		if lose[k] == 1 {
			answer1 = append(answer1, k)
		}
	}
	sort.Ints(answer0)
	sort.Ints(answer1)
	return [][]int{answer0, answer1}
}
