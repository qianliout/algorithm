package main

import "fmt"

func main() {
	fmt.Println("hello word")
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	dis := make([]int, n)
	inf := -(1 << 32)
	for i := range dis {
		dis[i] = inf
	}
	dis[0] = 0
	q := []int{0}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			if dis[y] >= 0 {
				continue
			}
			dis[y] = dis[x] + 1
			q = append(q, y)
		}
	}
	ans := 0
	for i := 1; i < n; i++ {
		d := dis[i] * 2
		p := patience[i]
		if p > 0 {
			ans = max(ans, (d-1)/p*p+d)
		}
	}
	return ans + 1
}
