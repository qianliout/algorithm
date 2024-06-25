package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(2, 91, 996))
}

func powerfulIntegers(x int, y int, bound int) []int {
	ans := make(map[int]bool)

	c, d := 1, 1
	for a := 0; a < 21; a++ {
		if c > bound || c <= 0 {
			break
		}
		d = 1
		for b := 0; b < 21; b++ {
			if c+d > bound || c+d <= 0 {
				break
			}
			ans[c+d] = true
			d = d * y
		}

		c = c * x
	}
	res := make([]int, 0)
	for k := range ans {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}
