package main

import (
	"sort"
)

func main() {

}

func sequentialDigits(low int, high int) []int {
	ans := make([]int, 0)
	for i := 1; i <= 9; i++ {
		ans = append(ans, help(low, high, i)...)
	}
	sort.Ints(ans)
	return ans
}

func help(low, high, start int) []int {
	num := start

	ans := make([]int, 0)
	for num <= high {
		if num >= low {
			ans = append(ans, num)
		}
		start++
		if start >= 10 {
			break
		}
		num = num*10 + start
	}

	return ans
}
