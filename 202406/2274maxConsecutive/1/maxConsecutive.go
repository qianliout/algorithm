package main

import (
	"fmt"
)

func main() {
	fmt.Print(maxConsecutive(2, 9, []int{4, 6}))
}

func maxConsecutive(bottom int, top int, special []int) int {
	ans := 0
	start := bottom
	spe := make(map[int]bool)
	for _, ch := range special {
		spe[ch] = true
	}
	for i := bottom; i <= top; i++ {
		if spe[i] {
			ans = max(ans, i-start)
			start = i + 1
		}
	}

	if !spe[top] {
		ans = max(ans, top-start+1)
	}

	return ans
}
