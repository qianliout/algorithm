package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(mostVisited(4, []int{1, 3, 1, 2}))
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
	fmt.Println(mostVisited(7, []int{1, 3, 5, 7}))
}

func mostVisited(n int, rounds []int) []int {
	for i := range rounds {
		rounds[i]--
	}

	ans := make([]int, n)
	ans[rounds[0]]++
	for i := 0; i < len(rounds)-1; i++ {
		start := rounds[i]
		end := rounds[i+1]
		if end <= start {
			end += n
		}
		for j := start + 1; j <= end; j++ {
			ans[j%n]++
		}
	}

	mx := slices.Max(ans)
	res := make([]int, 0)
	for i, ch := range ans {
		if ch == mx {
			res = append(res, i+1)
		}
	}
	return res
}
