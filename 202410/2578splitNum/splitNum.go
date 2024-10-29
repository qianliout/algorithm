package main

import (
	"sort"
)

func splitNum(num int) int {
	s := make([]int, 0)
	for num > 0 {
		s = append(s, num%10)
		num /= 10
	}
	sort.Ints(s)
	i := 0
	for i < len(s) {
		if s[i] != 0 {
			break
		}
		i++
	}
	s = s[i:]
	n := len(s)
	if n == 0 {
		return 0
	}
	a, b := 0, 0
	for j := 0; j+1 < n; j += 2 {
		a = a*10 + s[j]
		b = b*10 + s[j+1]
	}
	if n%2 == 1 {
		return min(a*10+s[n-1]+b, a+b*10+s[n-1])
	}

	return a + b
}
