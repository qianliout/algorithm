package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minMoves(5, 0))
	fmt.Println(minMoves(19, 2))
	fmt.Println(minMoves(5, 0))
}

// 1 <= target <= 109
// 数据范围太大，记忆化搜索是不行的
func minMoves1(target int, maxDoubles int) int {
	var dfs func(n int, us int) int
	inf := math.MaxInt >> 1
	dfs = func(n int, us int) int {
		if n > target || us > maxDoubles {
			return inf
		}
		if n == target {
			return 0
		}
		a := dfs(n+1, us) + 1
		b := dfs(n*2, us+1) + 1
		return min(a, b)
	}
	return dfs(1, 0)
}

// 反向会出错
func minMoves2(target int, maxDoubles int) int {
	if target == 1 {
		return 0
	}
	if target == 2 {
		return 1
	}
	if maxDoubles > 0 {
		a := target / 2
		b := target - a
		return minMoves(b, maxDoubles-1) + 1
	} else {
		return target - 1
	}
}

// 错误
func minMoves3(target int, maxDoubles int) int {
	var dfs func(n, us int) int
	inf := math.MaxInt >> 1
	dfs = func(n, us int) int {
		if n == target {
			return 0
		}
		if n > target || us > maxDoubles {
			return inf
		}
		if n*2 <= target && us < maxDoubles {
			return dfs(n*2, us+1) + 1
		} else {
			return dfs(n+1, us) + 1
		}
	}
	return dfs(1, 0)
}

func minMoves(target int, maxDoubles int) int {
	ans := 0
	for target > 1 {
		if maxDoubles <= 0 {
			return target - 1 + ans
		}
		if target&1 == 1 {
			target--
			ans++
		}
		target >>= 1
		ans++
		maxDoubles--
	}
	return ans
}
