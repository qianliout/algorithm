package main

import (
	"sort"
)

func main() {

}

func minimumSum(num int) int {
	ans := make([]int, 0)
	for num > 0 {
		ans = append(ans, num%10)
		num /= 10
	}
	sort.Ints(ans)
	a, b := 0, 0
	for i := 0; i < len(ans); i = i + 2 {
		a = a*10 + ans[i]
		b = b*10 + ans[i+1]
	}
	return a + b
}
