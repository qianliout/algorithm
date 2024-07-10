package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

// 对角线遍历
func findDiagonalOrder(nums [][]int) []int {
	n := len(nums)
	m := 0
	for _, ch := range nums {
		m = max(m, len(ch))
	}

	line := make([][]int, n+m-1)
	for i, ch := range nums {
		for j, ch2 := range ch {
			line[i+j] = append(line[i+j], ch2)
		}
	}
	ans := make([]int, 0)
	for i := range line {
		ch := line[i]
		reverse(ch)
		ans = append(ans, ch...)
	}

	return ans
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
