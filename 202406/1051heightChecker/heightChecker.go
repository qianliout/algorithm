package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

func heightChecker(heights []int) int {
	arr := append([]int{}, heights...)
	sort.Ints(arr)
	cnt := 0
	for i := 0; i < len(heights); i++ {
		if arr[i] != heights[i] {
			cnt++
		}
	}
	return cnt
}
