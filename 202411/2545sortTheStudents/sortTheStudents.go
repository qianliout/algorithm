package main

import (
	"sort"
)

func main() {

}

func sortTheStudents(score [][]int, k int) [][]int {
	if len(score) == 0 || len(score[0]) == 0 || k >= len(score[0]) {
		return score
	}
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}
