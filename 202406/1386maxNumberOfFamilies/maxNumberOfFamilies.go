package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(maxNumberOfFamilies(3, [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}))
	fmt.Println(maxNumberOfFamilies(4, [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}))

}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	res := make([]int, n+1)
	for _, ch := range reservedSeats {
		x, y := ch[0], ch[1]
		if y == 1 || y == 10 {
			continue
		}
		res[x] = res[x] | (1 << (9 - y))
	}
	cnt := n * 2
	for _, v := range res {
		cnt -= check(v)
	}
	return cnt

}

func check(set int) int {
	a := 0b000111100

	b := 0b11110000
	c := 0b00001111
	d := 0b11111111

	if set&d == 0 {
		return 2
	}
	return max(bti(set&a == 0), bti(set&b == 0), bti(set&c == 0))
}

func bti(a bool) int {
	if a {
		return 1
	}
	return 0
}

// 可以得到答案,但是会超内存，因为 n 会很大
func maxNumberOfFamilies1(n int, reservedSeats [][]int) int {

	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, 11)
	}
	for _, ch := range reservedSeats {
		x, y := ch[0], ch[1]

		matrix[x][y] = 1
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans += check1(matrix[i])
	}
	return ans
}

func check1(row []int) int {

	a := 0
	if !slices.Contains(row[4:8], 1) {
		a = 1
	}
	b := 0
	if !slices.Contains(row[2:6], 1) {
		b = 1
	}
	c := 0
	if !slices.Contains(row[6:10], 1) {
		c = 1
	}
	if b == 1 && c == 1 {
		return b + c
	}
	return max(a, b, c)
}

// 这样都还是超时
func maxNumberOfFamilies2(n int, reservedSeats [][]int) int {
	sort.Slice(reservedSeats, func(i, j int) bool {
		return reservedSeats[i][0] < reservedSeats[j][0]
	})
	ans := 0
	start := 0
	// se := len(reservedSeats)
	for i := 1; i <= n; i++ {
		j := start
		for j < len(reservedSeats) {
			if reservedSeats[j][0] != i {
				break
			}
			j++
		}
		ans += check2(reservedSeats[start:j])
		start = j
	}
	return ans
}

func check2(set [][]int) int {
	if len(set) == 0 {
		return 2
	}
	row := make([]int, 11)
	for _, ch := range set {
		row[ch[1]] = 1
	}

	a := 0
	if !slices.Contains(row[4:8], 1) {
		a = 1
	}
	b := 0
	if !slices.Contains(row[2:6], 1) {
		b = 1
	}
	c := 0
	if !slices.Contains(row[6:10], 1) {
		c = 1
	}
	if b == 1 && c == 1 {
		return b + c
	}
	return max(a, b, c)
}
