package main

import (
	"math/bits"
)

func main() {

}

func minimizeXor(num1 int, num2 int) int {
	a, b := bits.OnesCount(uint(num1)), bits.OnesCount(uint(num2))
	for a > b {
		num1 = num1 & (num1 - 1)
		b++
	}
	for a < b {
		num1 |= num1 + 1
		b--
	}
	return num1
}
