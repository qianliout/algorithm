package main

import (
	"fmt"
)

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1})) // [18,6,6,6,1,-1]
}

func replaceElements(arr []int) []int {
	n := len(arr)
	nums := make([]int, n)
	mx := -1
	for i := n - 1; i >= 0; i-- {
		nums[i] = mx
		mx = max(mx, arr[i])
	}
	return nums
}
