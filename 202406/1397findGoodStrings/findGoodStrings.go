package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findGoodStrings(2, "aa", "da", "b"))
	fmt.Println(findGoodStrings(8, "pzdanyao", "wgpmtywi", "sdka"))
}

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	mod := int(math.Pow10(9)) + 7

	var dfs func(i int, pre []byte, left, right bool) int

	// 这样写 mem 结果是错的
	mem := make([][]int, 4)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i int, pre []byte, left, right bool) int {
		if i >= n {
			if len(pre) > 0 {
				return 1
			}
			return 0
		}

		if va := mem[b2i(left, right)][i]; va != -1 {
			return va
		}

		low := byte('a')
		if left {
			low = byte(s1[i])
		}
		up := byte('z')
		if right {
			up = byte(s2[i])
		}
		res := 0
		for d := low; d <= up; d++ {
			ne := append([]byte{}, pre...)
			ne = append(ne, d)
			if sub(ne, evil) {
				continue
			}
			nex := dfs(i+1, ne, left && d == low, right && d == up)
			res = (res + nex) % mod
		}
		mem[b2i(left, right)][i] = res % mod

		return res % mod
	}
	return dfs(0, []byte{}, true, true) % mod
}

// 判断 b 是否是 a 的子串
func sub(a []byte, b string) bool {
	if len(a) < len(b) {
		return false
	}
	n := len(b)
	for i := 0; i+n <= len(a); i++ {
		if string(a[i:i+n]) == b {
			return true
		}
	}
	return false
}

func b2i(isLimit, isNum bool) int {
	if !isLimit && !isNum {
		return 0
	}
	if isNum && isLimit {
		return 1
	}
	if isNum {
		return 2
	}
	if isLimit {
		return 3
	}
	return 0
}
