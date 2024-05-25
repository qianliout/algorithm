package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
}

func mostCompetitive1(nums []int, k int) []int {
	win := make([]int, 0)
	n := len(nums)

	for i, ch := range nums {
		for len(win) > 0 && win[len(win)-1] > ch && (len(win)-1)+(n-i) >= k {
			win = win[:len(win)-1]
		}
		win = append(win, ch)
	}

	return win[:k]
}

func mostCompetitive(nums []int, k int) []int {
	win := make([]int, 0)
	n := len(nums)

	for i, ch := range nums {
		for len(win) > 0 && win[len(win)-1] > ch && (len(win)-1)+(n-i) >= k {
			win = win[:len(win)-1]
		}
		if len(win) < k {
			win = append(win, ch)
		}
	}

	return win
}
