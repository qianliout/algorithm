package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies(7, 4))
}

func distributeCandies(candies int, n int) []int {

	ans := make([]int, n)
	start := 0
	for ; candies > 0; start++ {
		idx := start % n
		ans[idx] += min(start+1, candies)
		candies -= min(start+1, candies)
	}
	return ans
}
