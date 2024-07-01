package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countTexts("222222222222222222222222222222222222"))
	fmt.Println(countTexts("55555555999977779555"))
}

var f []int
var g []int

var mod = int(math.Pow10(9)) + 7

func init() {
	mx := int(math.Pow10(5)) + 1
	f = make([]int, mx) // 2,3,4,5,6,8
	g = make([]int, mx) // 7,9
	f[0], f[1], f[2], f[3] = 1, 1, 2, 4
	g[0], g[1], g[2], g[3] = 1, 1, 2, 4
	for i := 4; i < mx; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

func countTexts(pressedKeys string) int {
	cnt := 0
	ans := 1
	n := len(pressedKeys)
	for i, ch := range pressedKeys {
		cnt++

		if i == n-1 || pressedKeys[i+1] != uint8(ch) {
			if ch == '7' || ch == '9' {
				ans = (ans * g[cnt]) % mod
			} else {
				ans = (ans * f[cnt]) % mod
			}
			cnt = 0
		}
	}

	return ans
}
