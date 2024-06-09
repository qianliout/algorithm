package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(minimumTime("1100101"))
	fmt.Println(minimumTime("0010"))
	fmt.Println(minimumTime("011001111111101001010000001010011"))
	fmt.Println(minimumTime("110001110000100001100010111101010011101101000111"))
}

// 可以得到答案，但是会超时
func minimumTime(s string) int {
	mem := make(map[string]int)
	return dfs(s, mem)
}

func dfs(ss string, mem map[string]int) int {
	s := []byte(ss)
	if len(s) == 0 {
		return 0
	}
	k := strings.Count(ss, "1")
	if k == 0 {
		mem[ss] = 0
		return 0
	}
	if va, ok := mem[ss]; ok {
		return va
	}

	n := len(s)
	res := math.MaxInt
	res = min(res, dfs(string(s[1:]), mem)+1)
	res = min(res, dfs(string(s[:n-1]), mem)+1)

	for i := 1; i < n-1; i++ {
		ns := append([]byte{}, s[:i]...)
		ns = append(ns, s[i+1:]...)
		res = min(res, dfs(string(ns), mem)+2)
	}
	mem[ss] = res
	return res
}
