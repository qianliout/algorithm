package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(stringCount(4))
}

// 正难则反
/*
不含 leet 的字符串，需要至少满足如下三个条件中的一个：
    不含 l。
    不含 t。
    不含 e 或者恰好包含一个 e。
*/

// 不含 leet 的字符串的个数为「至少满足一个条件」减去「至少满足两个条件」加上「满足三个条件」，这就是容斥原理
func stringCount2(n int) int {
	mod := int(math.Pow10(9)) + 7
	return ((pow(26, n, mod)-
		pow(25, n-1, mod)*(75+n)+
		pow(24, n-1, mod)*(72+n*2)-
		pow(23, n-1, mod)*(23+n))%mod + mod) % mod // 保证结果非负
}

func stringCount(n int) int {
	mod := int(math.Pow10(9)) + 7
	all := pow(26, n, mod)
	n = n - 1
	// 不满足一个条件
	a := pow(25, n, mod)*2 + pow(25, n, mod) + n*pow(24, n-1, mod)
	b := pow(24, n, mod) + (pow(24, n, mod)+n*pow(24, n-1, mod))*2
	c := pow(23, n, mod) + n*pow(23, n-1, mod)

	return ((all-a-b+c)%mod + mod) % mod

}
func pow(x, n int, mod int) int {
	res := 1
	for ; n > 0; n = n / 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
