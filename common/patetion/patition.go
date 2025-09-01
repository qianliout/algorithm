package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(Patition(10, 5))
	fmt.Println(Patition(10, 7))
	fmt.Println(Patition(20, 7))
	fmt.Println(Patition(30, 1))
	fmt.Println(Patition(30, 2))
}

func Patition(m int, n int) []float64 {
	ans := patition(m*100, n)
	res := make([]float64, n)
	for i := range res {
		res[i] = float64(ans[i]) / 100
	}
	return res
}

func patition(all int, n int) []int {
	ans := make([]int, 0)
	var dfs func(m int, path []int, mx, mi int)
	dfs = func(m int, path []int, mx, mi int) {
		if len(path) == n {
			if m == 0 && mx < mi*2 {
				ans = append(ans, path...)
				return
			}
			return
		}

		if len(ans) == n {
			return
		}

		if m < n-len(path) {
			return
		}
		i := rand.Intn(m) + 1
		path = append(path, i)
		dfs(m-i, path, max(mx, i), min(mi, i))
		path = path[:len(path)-1]
	}

	for {
		dfs(all, []int{}, 0, all+1)
		if len(ans) == n {
			return ans
		}
	}
}

func DistributeCents(m, n int) []float64 {
	mi := int(math.Floor(float64(m*100) / float64(n+1)))
	mx := int(math.Ceil(float64(2*m*100) / float64(2*n-1)))
	ans := make([]int, n)
	for i := range ans {
		ans[i] = rand.Intn(mx-mi) + 1
	}
	// 归一化
	
}
