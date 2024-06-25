package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{59, 91}, 50))
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}

func bagOfTokensScore(tokens []int, power int) int {
	n := len(tokens)
	sort.Ints(tokens)
	ans := 0
	le, ri := 0, n-1
	score := 0
	for le <= ri {
		if power >= tokens[le] {
			score++
			ans = max(ans, score)
			power -= tokens[le]
			le++
		} else if ans > 0 {
			power += tokens[ri]
			score--
			ans = max(ans, score)
			ri--
		} else {
			break
		}
	}
	return ans
}
