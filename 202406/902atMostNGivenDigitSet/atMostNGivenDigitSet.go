package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000))
}

func atMostNGivenDigitSet(digits []string, n int) int {
	var dfs func(i int, isLimit, isNum bool) int
	s := strconv.Itoa(n)
	mem := make([][]int, 4)
	for i := range mem {
		mem[i] = make([]int, len(s))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i int, isLimit, isNum bool) int {
		if i >= len(s) {
			if isNum {
				return 1
			}
			return 0
		}
		res := 0
		if va := mem[b2i(isLimit, isNum)][i]; va != -1 {
			return va
		}
		if !isNum {
			res += dfs(i+1, false, false)
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for _, d1 := range digits {
			d, _ := strconv.Atoi(d1)
			// 这里就限制了，数字不会超过 n
			if d > up {
				// 因为是非递减的
				break
			}
			res += dfs(i+1, isLimit && d == up, true)
		}
		mem[b2i(isLimit, isNum)][i] = res
		return res
	}
	return dfs(0, true, false)
}

func b2i(isLimit, isNum bool) int {
	if !isLimit && !isNum {
		return 0
	}
	if isNum && isLimit {
		return 1
	}
	if isNum {
		return 2
	}
	if isLimit {
		return 3
	}
	return 0
}
