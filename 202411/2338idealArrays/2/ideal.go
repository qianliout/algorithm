package main

import (
	"fmt"
)

func main() {
	fmt.Println(idealArrays(5, 9))
}

const mod, mx, mxK int = 1e9 + 7, 1e4 + 1, 13 // 至多 13 个质因数

var ks [mx][]int

var c [mx + mxK][mxK + 1]int

func init() {
	for i := 2; i < mx; i++ {
		x := i
		for p := 2; p*p <= x; p++ {
			if x%p == 0 {
				k := 1
				for x /= p; x%p == 0; x /= p {
					k++
				}
				ks[i] = append(ks[i], k)
			}
		}
		if x > 1 {
			ks[i] = append(ks[i], 1)
		}
	}
	c[0][0] = 1
	for i := 1; i < len(c); i++ {
		c[i][0] = 1
		for j := 1; j <= mxK && j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}
}

func idealArrays(n, maxValue int) (ans int) {
	for _, kss := range ks[1 : maxValue+1] {
		mul := 1
		for _, k := range kss {
			mul = mul * c[n+k-1][k] % mod
		}
		ans = (ans + mul) % mod
		fmt.Println("ans ", ans)
	}
	return ans
}
