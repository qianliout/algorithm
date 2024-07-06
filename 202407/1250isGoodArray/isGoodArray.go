package main

import "fmt"

func main() {
	fmt.Println(isGoodArray([]int{6, 10, 15}))
}

func isGoodArray(nums []int) bool {
	g := 0
	for _, x := range nums {
		g = gcb(g, x)
	}

	return g == 1
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}
