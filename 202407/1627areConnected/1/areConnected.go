package main

import (
	"fmt"

	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(areConnected(6, 2, [][]int{{1, 4}, {2, 5}, {3, 6}}))
}

var cnt [][]int

func init() {
	cnt = make([][]int, 10010)
	for i := range cnt {
		cnt[i] = make([]int, 10010)
	}
}

// 会超内存
func areConnected(n int, threshold int, queries [][]int) []bool {

	ans := make([]bool, len(queries))
	if threshold <= 0 {
		for i := range ans {
			ans[i] = true
		}
		return ans
	}

	uf := NewSizeUnionFind(n + 2)
	for i := 1; i <= n; i++ {
		// j:=i,不能是j:=i+1,因为threshold可能是1
		for j := i; j <= n; j++ {
			l := gcb(i, j)
			if l > threshold {
				uf.Union(i, j)
			}
		}
	}

	for i, ch := range queries {
		x, y := ch[0], ch[1]
		ans[i] = uf.IsConnected(x, y)
	}
	return ans
}

func lcm(a, b int) int {
	return a * b / gcb(a, b)
}

func gcb(a, b int) int {
	if cnt[a][b] > 0 {
		return cnt[a][b]
	}
	if b == 0 {
		cnt[a][b] = a
		return a
	}
	c := gcb(b, a%b)
	cnt[a][b] = c
	return c
}
