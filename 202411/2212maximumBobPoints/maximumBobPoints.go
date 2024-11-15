package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}))
	fmt.Println(maximumBobPoints(3, []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}))
	fmt.Println(maximumBobPoints(86937, []int{0, 4196, 17248, 4008, 23425, 2037, 13917, 1227, 1623, 11693, 1816, 5747}))
}

// 当numArrows 很大时会超时
func maximumBobPoints2(numArrows int, aliceArrows []int) []int {
	n := len(aliceArrows)
	var dfs func(i, j int) int
	res := make([]int, n)
	// inf := 1 << 32
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, numArrows+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if j <= 0 || i < 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		ans := 0
		for k := j; k >= 0; k-- {
			if aliceArrows[i] < k {
				ans = max(ans, dfs(i-1, j-k)+i)
			} else {
				ans = max(ans, dfs(i-1, j-k))
			}
		}
		mem[i][j] = ans
		return ans
	}
	dfs(n-1, numArrows)
	for i := 11; i > 0; i-- {
		if mem[i][numArrows] > mem[i-1][numArrows] {
			aa := aliceArrows[i]
			res[i] = aa + 1
			numArrows -= aa + 1
		}
	}
	res[0] = numArrows

	return res
}

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	n := len(aliceArrows)
	res := make([]int, n)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, numArrows+1)
	}
	for i := 1; i <= 11; i++ {
		aa := aliceArrows[i]
		for j := 1; j <= numArrows; j++ {
			if j < aa+1 {
				mem[i][j] = mem[i-1][j]
			} else {
				mem[i][j] = max(mem[i-1][j-aa-1]+i, mem[i-1][j])
			}
		}
	}

	// 状态恢复
	for i := 11; i > 0; i-- {
		if mem[i][numArrows] > mem[i-1][numArrows] {
			aa := aliceArrows[i]
			res[i] = aa + 1
			numArrows -= aa + 1
		}
	}
	res[0] = numArrows

	return res
}
