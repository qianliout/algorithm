package main

import (
	"fmt"
	"math/bits"
)

func main() {

	fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))

}

func maximumRows(matrix [][]int, numSelect int) int {
	m := len(matrix)
	n := len(matrix[0])
	g := make([]int, m)
	for i := 0; i < m; i++ {
		g[i] = cal(matrix[i])
	}
	ans := 0
	for i := 0; i < 1<<n; i++ {
		if bits.OnesCount64(uint64(i)) != numSelect {
			continue
		}
		ans = max(ans, count(g, i))
	}
	return ans
}

func cal(a []int) int {
	ans := 0
	m := len(a)
	for j, c := range a {
		ans = ans | (c << (m - j - 1))
	}
	return ans
}

func count(g []int, a int) int {
	ans := 0
	for _, c := range g {
		if c&a <= a {
			ans++
		}
	}
	return ans
}
