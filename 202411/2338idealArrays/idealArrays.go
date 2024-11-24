package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(idealArrays(2, 5))
	// fmt.Println(idealArrays(5, 9))
	// fmt.Println(idealArrays(184, 389))
	// fmt.Println(idealArrays(380, 194))
	fmt.Println(idealArrays(5878, 2900))
	// // fmt.Println(CalPrime2(9))
	// fmt.Println(CalPrime(9))
	// fmt.Println(CalPrime(8))
	// fmt.Println(CalPrime(7))
	// fmt.Println(CalPrime(6))

}

func idealArrays(n int, maxValue int) int {
	// 分解质因子
	mod := int(math.Pow10(9)) + 7
	cnt := 13 // 最多有13个质因数
	com := Comb(n+cnt, cnt, mod)
	ans := 0
	for i := 1; i <= maxValue; i++ {
		mul := 1
		ks := CalPrime(i)
		for _, k := range ks {
			mul = mul * com[n+k-1][k] % mod
		}
		ans += mul
	}

	return ans % mod
}

// CalPrime 求每个质数的个数
func CalPrime(n int) map[int]int {
	ks := make(map[int]int)
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			for n%p == 0 {
				ks[p]++
				n = n / p
			}
		}
	}
	if n > 1 {
		ks[n] = 1
	}
	return ks
}

// Comb 求组合数C(n,m) 他可以用递推得到 C(n,m) = C(n-1,m-1) + C(n-1,m)
func Comb(n, k int, mod int) [][]int {
	ans := make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, k+1)
	}
	// 初值
	ans[0][0] = 1
	for i := 1; i <= n; i++ {
		ans[i][0] = 1
		for j := 1; j <= k && j <= i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			// 这里一定要取模，不然就会有溢出
			ans[i][j] = ans[i][j] % mod
		}
	}
	return ans
}
