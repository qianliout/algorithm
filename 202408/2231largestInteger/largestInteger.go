package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestInteger(1234))
	fmt.Println(largestInteger(247))
}

func largestInteger(num int) int {
	bi := make([]int, 0)
	for num > 0 {
		bi = append(bi, num%10)
		num /= 10
	}
	l, r := 0, len(bi)-1
	for l < r {
		bi[l], bi[r] = bi[r], bi[l]
		l++
		r--
	}
	// 使用原地排序的思想
	n := len(bi)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if bi[i] < bi[j] && bi[i]&1 == bi[j]&1 {
				bi[i], bi[j] = bi[j], bi[i]
			}
		}
	}
	ans := 0
	for _, ch := range bi {
		ans = ans*10 + ch
	}
	return ans
}
