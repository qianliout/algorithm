package main

import (
	"math"
)

func main() {

}

func closestDivisors(num int) []int {
	ans1 := help(num + 1)
	ans2 := help(num + 2)
	if abs(ans1[0]-ans1[1]) < abs(ans2[0]-ans2[1]) {
		return ans1
	}
	return ans2
}

func help(num int) []int {
	ans := make([]int, 0)
	res := divis(num)
	for _, ch := range res {
		if len(ans) == 0 || abs(ans[0]-ans[1]) > abs(ch.b-ch.a) {
			ans = []int{ch.a, ch.b}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func divis(num int) []pair {
	ans := make([]pair, 0)
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			ans = append(ans, pair{i, num / i})
		}
	}
	return ans
}

type pair struct {
	a, b int
}
