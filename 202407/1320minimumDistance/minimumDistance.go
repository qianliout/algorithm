package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDistance("CAKE"))
}

func minimumDistance(word string) int {
	inf := math.MaxInt / 100
	n := len(word)
	state := make([][][]int, n+1)
	for i := range state {
		state[i] = make([][]int, 26)
		for j := range state[i] {
			state[i][j] = make([]int, 26)
		}
	}
	// 这里 i=0不能 赋值，不然得不到正确的结果
	for i := 1; i <= n; i++ {
		for j := range state[i] {
			for k := range state[i][j] {
				state[i][j][k] = inf
			}
		}
	}

	ans := inf
	for i := 1; i <= n; i++ {
		v := int(word[i-1] - 'A')
		for l := 0; l < 26; l++ {
			for r := 0; r < 26; r++ {
				// 判断上一个阶段的状态是否存在
				if state[i-1][l][r] != inf {
					// 移动左指
					state[i][v][r] = min(state[i][v][r], state[i-1][l][r]+help(l, v))
					// 移动右指
					state[i][l][v] = min(state[i][l][v], state[i-1][l][r]+help(r, v))
				}
				if i == n {
					ans = min(ans, state[i][v][r])
					ans = min(ans, state[i][l][v])
				}
			}
		}
	}
	return ans
}

func help(x, y int) int {
	x0, y0 := x/6, x%6
	x1, y1 := y/6, y%6
	return abs(x1-x0) + abs(y0-y1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
