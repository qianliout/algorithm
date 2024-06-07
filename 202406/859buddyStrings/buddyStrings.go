package main

import (
	"fmt"
)

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}

func buddyStrings(s string, goal string) bool {
	cnt1 := 0
	cnt2 := 0
	if len(s) != len(goal) {
		return false
	}
	exit1 := make(map[uint8]int)
	exit2 := make(map[uint8]int)
	for i := range s {
		exit1[s[i]]++
		exit2[goal[i]]++
		cnt2 = max(exit1[s[i]], cnt2)
		if s[i] != goal[i] {
			cnt1++
		}
	}
	if !same(exit1, exit2) {
		return false
	}
	if cnt1 == 2 || cnt1 == 0 && cnt2 >= 2 {
		return true
	}
	return false
}

func same(a, b map[uint8]int) bool {

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, ch := range bills {
		switch ch {
		case 5:
			five++
		case 10:
			if five <= 0 {
				return false
			}
			five--
			ten++
		case 20:
			sub := 15
			if ten > 0 {
				ten--
				sub -= 10
			}
			for sub > 0 && five > 0 {
				sub -= 5
				five--
			}
			if sub > 0 {
				return false
			}
		}
	}
	return true
}

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	col, row := len(matrix), len(matrix[0])
	ans := make([][]int, row)
	for i := range ans {
		ans[i] = make([]int, col)
	}
	for i := range matrix {
		for j, ch := range matrix[i] {
			ans[j][i] = ch
		}
	}
	return ans
}
