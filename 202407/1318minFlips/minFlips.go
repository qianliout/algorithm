package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlips(2, 6, 5))
}

func minFlips(a int, b int, c int) int {
	cnt := 0
	for {
		d := a & 1
		e := b & 1
		f := c & 1
		if f == 0 {
			cnt += d + e
		}
		if f == 1 {
			if d == 0 && e == 0 {
				cnt++
			}
		}
		a >>= 1
		b >>= 1
		c >>= 1
		if a == 0 && b == 0 && c == 0 {
			break
		}

	}
	return cnt
}
