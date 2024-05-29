package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkRecord(2))
}

func checkRecord(n int) int {
	mem := make([][][]int, n+1)
	for i := range mem {
		mem[i] = make([][]int, 2)
		for j := range mem[i] {
			mem[i][j] = make([]int, 3)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}
	return dfd(n, 0, 0, mem)
}

func dfd(n int, ac int, lc int, mem [][][]int) int {
	// 这里的二个判断，前面判断一定在后面的前面，不然结果会错
	if ac >= 2 || lc >= 3 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if mem[n][ac][lc] >= 0 {
		return mem[n][ac][lc]
	}

	base := int(math.Pow(10, 9)) + 7

	ans := 0
	ans = (ans + dfd(n-1, ac+1, 0, mem)) % base  // A
	ans = (ans + dfd(n-1, ac, lc+1, mem)) % base // L
	ans = (ans + dfd(n-1, ac, 0, mem)) % base    // P
	mem[n][ac][lc] = ans
	return ans
}
