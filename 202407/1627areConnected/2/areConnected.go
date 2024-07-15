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

func areConnected(n int, threshold int, queries [][]int) []bool {

	ans := make([]bool, len(queries))

	uf := NewSizeUnionFind(n + 2)
	// 枚举公因数
	for z := threshold + 1; z <= n; z++ {
		// 枚举两个 z 的倍数的点并连接
		for p, q := z, z*2; q <= n; p, q = p+z, q+z {
			uf.Union(p, q)
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
