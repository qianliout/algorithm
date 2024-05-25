package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrangeBookshelf([]int{5, 5, 6, 5}, 3))
	fmt.Println(arrangeBookshelf([]int{3, 3, 9, 8, 9, 2, 8}, 1))
	fmt.Println(arrangeBookshelf([]int{10, 4, 12, 1, 6, 10, 1, 10, 2, 10, 10, 7}, 1))
}

func arrangeBookshelf(order []int, limit int) []int {
	last := make(map[int]int)
	for _, ch := range order {
		last[ch]++
	}
	add := make(map[int]int)
	win := make([]int, 0)
	for _, ch := range order {
		if add[ch] >= limit {
			last[ch]--
			continue
		}
		for len(win) > 0 && win[len(win)-1] > ch && (add[win[len(win)-1]]+last[win[len(win)-1]] > limit) {
			add[win[len(win)-1]]--
			win = win[:len(win)-1]
		}
		win = append(win, ch)
		add[ch]++
		last[ch]--
	}
	return win
}
