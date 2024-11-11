package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))
}
func makeIntegerBeautiful(n int64, target int) int64 {
	var mul int64 = 1
	for {
		m := n + (mul-n%mul)%mul
		x := m
		sum := 0
		for x > 0 {
			sum += int(x % 10)
			x = x / 10
		}
		if sum <= target {
			return m - n
		}
		mul = mul * 10
	}
}
