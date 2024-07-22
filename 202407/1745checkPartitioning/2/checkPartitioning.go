package main

import (
	"fmt"
)

func main() {
	// fmt.Println(checkPartitioning("aba"))
	fmt.Println(checkPartitioning("abcbdd"))
	fmt.Println(checkPartitioning("acab"))
}

func checkPartitioning(s string) bool {
	n := len(s)
	check := make([][]bool, n)
	for i := range check {
		check[i] = make([]bool, n)
	}
	// 预处理回文串 i向右 j=i j向左遍历
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if i == j {
				check[j][i] = true
				continue
			}
			// 这样判断会有错
			if s[i] == s[j] && j+1 <= i-1 && check[j+1][i-1] {
				check[j][i] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+1 > i || i+1 > n-1 {
				continue
			}
			if check[0][j] && check[j+1][i] && check[i+1][n-1] {
				return true
			}
		}
	}

	return false
}
