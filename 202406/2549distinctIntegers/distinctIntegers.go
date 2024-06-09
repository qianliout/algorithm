package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(monkeyMove(6))
	fmt.Println(monkeyMove(3))
	fmt.Println(monkeyMove(4))
}

func distinctIntegers(n int) int {
	// n %(n-1) =1 根据这个就能得出，所有的 n,都能得到 n-1

	return max(1, n-1)
}

func monkeyMove(n int) int {
	// 总的次数，减去不会发生碰撞的次数
	// 总的次数：每个猴子有两种方式，总的就是2的 n次方
	// 不会发生碰撞的次数：所有猴子同时顺时针走，或同时逆时针走，总共2 次
	mod := int(math.Pow10(9) + 7)
	ans := ((pow(2, n, mod)-2)%mod + mod) % mod

	return ans
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
