package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxConsecutive(2, 9, []int{4, 6}))
	fmt.Println(maxConsecutive(6, 8, []int{7, 6, 8}))
}

func maxConsecutive(bottom int, top int, special []int) int {
	special = append(special, bottom, top)
	sort.Ints(special)
	start := bottom
	ans := 0
	for i := 1; i < len(special); i++ {
		ans = max(ans, special[i]-start)
		if i == len(special)-1 {
			ans = max(ans, special[i]-start+1)
		}
		start = special[i] + 1
	}

	return ans
}
