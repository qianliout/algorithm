package main

import (
	"fmt"
)

func main() {
	fmt.Println(countOfPairs(3, 1, 3))
}

func countOfPairs(n int, x int, y int) []int {
	res := make([]int, n)
	for i := 1; i <= n; i++ {
		bsf(i, n, x, y, res)
	}
	return res
}

func bsf(start, n, x, y int, res []int) {
	visit := make([]bool, n+1)
	queue := make([]int, 0)
	queue = append(queue, start)
	dis := 0
	visit[start] = true
	for len(queue) > 0 {
		dis++
		lev := make([]int, 0)
		for _, num := range queue {
			adj := getAdjacent(num, n, x, y)
			for _, d := range adj {
				if visit[d] {
					continue
				}
				visit[d] = true
				lev = append(lev, d)
				res[dis-1]++
			}
		}
		queue = lev
	}
}

func getAdjacent(num, n, x, y int) []int {
	ans := make([]int, 0)
	if num > 1 {
		ans = append(ans, num-1)
	}
	if num < n {
		ans = append(ans, num+1)
	}
	if num == y && abs(x-y) > 1 {
		ans = append(ans, x)
	}
	if num == x && abs(x-y) > 1 {
		ans = append(ans, y)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
