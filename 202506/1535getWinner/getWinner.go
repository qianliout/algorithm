package main

import (
	"fmt"
)

func main() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
}

func getWinner(arr []int, k int) int {
	mx := arr[0]
	win := 0
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] > mx {
			mx = arr[i]
			win = 0
		}
		win++
		if win >= k {
			return mx
		}
	}
	return mx
}
