package main

import (
	"fmt"
)

func main() {

}

func numEquivDominoPairs(dominoes [][]int) int {
	for i, ch := range dominoes {
		x, y := ch[0], ch[1]
		dominoes[i] = []int{max(x, y), min(x, y)}
	}
	// sort.Slice(dominoes, func(i, j int) bool { return dominoes[i][0] <= dominoes[j][0] })

	cnt := make(map[string]int)
	for _, ch := range dominoes {
		key := fmt.Sprintf("%d-%d", ch[0], ch[1])
		cnt[key]++
	}
	ans := 0
	for _, v := range cnt {
		if v <= 1 {
			continue

		}
		ans += (v * (v - 1)) / 2
	}
	return ans

}
