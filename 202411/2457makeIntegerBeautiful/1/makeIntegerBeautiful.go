package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))
}

// 没有理解题目
func makeIntegerBeautiful(n int64, target int) int64 {
	ss := make([]int, 0)
	for n > 0 {
		ss = append(ss, int(n%10))
		n = n / 10
	}
	l, r := 0, len(ss)-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}
	sum := 0
	for _, ch := range ss {
		sum += ch
	}
	mul := 0
	i := len(ss) - 1
	var ans int64
	for i >= 0 {
		a := (9 - ss[i]) * mul
		if sum+a <= target {
			ans += int64(a)
			mul = mul * 10
			i--
			sum += a
		} else {
			ans += int64(target - sum)
			return ans
		}
	}
	return ans + int64(target-sum)
}
