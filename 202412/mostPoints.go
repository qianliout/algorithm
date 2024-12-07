package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	fmt.Println(mostPoints1([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
}

func mostPoints1(questions [][]int) int64 {
	n := len(questions)
	var dfs func(i int) int

	dfs = func(i int) int {
		if i < 0 || i >= n {
			return 0
		}
		c, y := questions[i][0], questions[i][1]
		a := c + dfs(i+y+1)
		b := dfs(i + 1)
		return max(a, b)
	}
	ans := dfs(0)
	return int64(ans)
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int64, n+10)

	for i := n - 1; i >= 0; i-- {
		c, y := questions[i][0], questions[i][1]
		f[i] = max(f[i+1], int64(c)+f[min(n, i+y+1)])
	}

	return f[0]
}
