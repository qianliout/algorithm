package main

import (
	"fmt"
)

func main() {
	fmt.Println(substringXorQueries("101101", [][]int{{0, 5}, {1, 2}}))
	fmt.Println(substringXorQueries("11111111011010011101000110010111110", [][]int{{4, 186}, {1, 2}}))

}

func substringXorQueries(s string, queries [][]int) [][]int {
	type pair struct{ le, ri int }
	n := len(s)
	// 预处理所有的数据，最大的数是10……9，最大是30位，所以预处理 s中1-30位内的数字
	mem := make(map[int]pair)
	for l := 0; l < n; l++ {
		x := 0
		for r := l; r < min(l+30, n); r++ {
			x = x<<1 | int(s[r]-'0')&1
			if va, ok := mem[x]; !ok || va.ri-va.le > r-l {
				mem[x] = pair{l, r}
			}
		}
	}
	ans := make([][]int, 0)
	for i := range queries {
		ch := queries[i]
		vl, ok := mem[ch[0]^ch[1]]
		if !ok {
			ans = append(ans, []int{-1, -1})
		} else {
			ans = append(ans, []int{vl.le, vl.ri})
		}
	}
	return ans
}
