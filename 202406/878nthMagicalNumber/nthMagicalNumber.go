package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthMagicalNumber(4, 2, 3))
	fmt.Println(nthMagicalNumber(1000000000, 40000, 40000))
}

/*
一个正整数如果能被 a 或 b 整除，那么它是神奇的。
给定三个整数 n , a , b ，返回第 n 个神奇的数字
*/
func nthMagicalNumber(n int, a int, b int) int {
	mi, mx := 0, min(a, b)*n+1
	// 二分找
	mod := int(math.Pow10(9)) + 7
	le, ri := mi, mx
	gb := gbc(a, b)
	for le < ri {
		// 左端点
		mid := le + (ri-le)/2
		c := check(mid, a, b, gb)
		if mid >= mi && mid < mx && c >= n {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le % mod
}

func check(mid, a, b, gb int) int {
	c := a / gb * b // 这个表示 a,b 的最小公倍数
	return mid/a + mid/b - mid/c
}

func gbc(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
