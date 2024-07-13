package main

import (
	"fmt"
)

func main() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
}

func getWinner(arr []int, k int) int {
	// 这个题目的难点就是首位值的处理
	mx := arr[0]
	win := -1
	for _, x := range arr {
		if x > mx {
			mx = x
			win = 0
		}
		win += 1
		if win == k {
			break
		}
	}
	return mx
}
