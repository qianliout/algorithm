package main

import (
	"sort"
)

func main() {

}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	ans, n, m := 0, len(players), len(trainers)
	i, j := 0, 0
	for i < n && j < m {
		if players[i] <= trainers[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return ans
}
