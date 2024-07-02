package main

import (
	"fmt"
	"strings"
)

func main() {
	v := strings.Repeat(string("UD"[1]), 2)
	fmt.Println(v)
	fmt.Println(alphabetBoardPath("leet")) // DDR!UURRR!!DDD!
}

func alphabetBoardPath(target string) string {
	ans := make([]string, 0)
	x, y := 0, 0
	for _, ch := range target {
		nx, ny := (int(ch)-'a')/5, (int(ch)-'a')%5
		// ,竖直
		// 如果 nx==x,也就是说有相同的数字，这里重复0次，所以不用特别判定
		v := strings.Repeat(b2i("UD", nx < x), abs(nx-x))
		// 水平
		h := strings.Repeat(b2i("LR", ny < y), abs(ny-y))
		if ch == 'z' {
			v, h = h, v
		}
		ans = append(ans, v+h+"!")
		x, y = nx, ny
	}
	return strings.Join(ans, "")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func b2i(st string, a bool) string {
	if a {
		return string(st[0])
	}
	return string(st[1])
}
