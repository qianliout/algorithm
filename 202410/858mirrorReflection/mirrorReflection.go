package main

import (
	"fmt"
)

func main() {
	fmt.Println(3 % 6)
}

func mirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p = (p / g) & 1
	q = (q / g) & 1
	if p == 1 && q == 1 {
		return 1

	}
	if p == 1 {
		return 0
	}

	return 2
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 递归的辗转相除法
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
