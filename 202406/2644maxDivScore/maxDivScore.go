package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDivScore([]int{2, 9, 15, 50}, []int{5, 3, 7, 2}))
}

func maxDivScore(nums []int, divisors []int) int {
	mx := -1
	num := 0
	for _, div := range divisors {
		res := 0
		for _, ch := range nums {
			if ch%div == 0 {
				res++
			}
		}
		if res > mx || (res == mx && div < num) {
			mx = res
			num = div
		}
	}
	return num
}

func rowAndMaximumOnes(mat [][]int) []int {
	n := len(mat)
	mx, ans := 0, make([]int, 1)
	for i := 0; i < n; i++ {
		s := 0
		for _, ch := range mat[i] {
			s += ch
		}
		if s > mx {
			mx = s
			ans = []int{i, s}
		}
	}
	return ans
}
