package main

import (
	"sort"
)

func main() {

}

// numsDivide 中每个数都能整除，意思就是可以整除numsDivide 的最大公约数
func minOperations(nums []int, numsDivide []int) int {
	g := 0
	for _, ch := range numsDivide {
		g = gcd(ch, g)
	}
	sort.Ints(nums)
	for i, ch := range nums {
		if g%ch == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
