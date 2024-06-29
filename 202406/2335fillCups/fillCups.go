package main

import (
	"sort"
)

func main() {

}

func fillCups(amount []int) int {
	sort.Ints(amount)
	a, b, c := amount[0], amount[1], amount[2]
	if a+b <= c {
		return c
	}
	t := a + b - c
	// 这里为啥是 t+1呢，除法会向下取整数，
	return (t+1)/2 + c
}
