package main

import (
	"fmt"
)

func main() {
	fmt.Println(minImpossibleOR([]int{2, 1}))
	fmt.Println(minImpossibleOR([]int{5, 3, 2}))
}

func minImpossibleOR(nums []int) int {
	exit := make(map[int]int)
	for _, ch := range nums {
		exit[ch] = 1
	}
	i := 0
	for i <= 63 {
		a := 1 << i
		if exit[a] <= 0 {
			return a
		}
		i++
	}

	return -1
}
