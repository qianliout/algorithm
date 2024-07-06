package main

import "fmt"

func main() {
	fmt.Println(reconstructMatrix(4, 7, []int{2, 1, 2, 2, 1, 1, 1}))
}

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	ans := make([][]int, 0)
	ans1 := make([]int, n)
	ans2 := make([]int, n)
	for i := 0; i < n; i++ {
		if colsum[i] == 2 {
			if upper == 0 || lower == 0 {
				return ans
			}
			ans1[i] = 1
			ans2[i] = 1
			upper--
			lower--
		}
		if colsum[i] == 1 {
			if upper == 0 && lower == 0 {
				return ans
			}
			if upper <= lower {
				ans2[i] = 1
				lower--
			} else {
				ans1[i] = 1
				upper--
			}
		}
	}
	if upper > 0 || lower > 0 {
		return ans
	}
	return [][]int{ans1, ans2}
}
