package main

import (
	"fmt"
)

func main() {
	fmt.Println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))
}

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	d := make([]int, n+1)
	for _, ch := range shifts {
		x, y, z := ch[0], ch[1], ch[2]
		if z == 0 {
			d[x] -= 1
			d[y+1] += 1
		} else if z == 1 {
			d[x] += 1
			d[y+1] -= 1
		}
	}
	ans := make([]byte, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += d[i]
		ans[i] = f(s[i], sum)
	}
	return string(ans)
}

func f(a byte, n int) byte {
	b := (((int(a-'a')+n)%26 + 26) % 26) + 'a'
	return byte(b)
}
