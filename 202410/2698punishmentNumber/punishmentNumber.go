package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(punishmentNumber(10))
	fmt.Println(punishmentNumber(37))
}

func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if ok[i] {
			ans = ans + i*i
		}
	}
	return ans
}

// 打表
var ok []bool

func init() {
	ok = make([]bool, 1010)
	for i := range ok {
		ok[i] = check(i, 0, help2(i*i))
	}
}

func check(n int, pre int, ss []byte) bool {
	if n == pre && len(ss) == 0 {
		return true
	}
	for i := 1; i <= len(ss); i++ {
		a := help1(ss[:i])
		pre = pre + a
		if check(n, pre, ss[i:]) {
			return true
		}
		pre = pre - a
	}
	return false
}

func help1(ss []byte) int {
	a := 0
	for _, ch := range ss {
		a = a*10 + int(ch) - int('0')
	}
	return a
}

func help2(n int) []byte {
	a := strconv.Itoa(n)
	return []byte(a)
}
