package main

import (
	"fmt"
)

func main() {
	fmt.Println(prevPermOpt1([]int{3, 1, 1, 3}))
}

func prevPermOpt1(arr []int) []int {
	left := -1
	n := len(arr)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			left = i
			break
		}
	}
	if left == -1 {
		return arr
	}
	mx := 0
	right := n
	for i := left + 1; i < n; i++ {
		if arr[i] < arr[left] && mx < arr[i] {
			mx = arr[i]
			right = i
			// break
		}
	}
	if right == n {
		return arr
	}
	arr[left], arr[right] = arr[right], arr[left]
	return arr
}
