package main

import (
	"fmt"
)

func main() {
	// fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("2*3+4"))
}

func calculate(s string) int {
	s = s + "+"
	op := "+"
	ss := []byte(s)
	st := make([]int, 0)
	num := 0
	for _, ch := range ss {
		if ch == ' ' {
			continue
		}
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}
		if op == "+" {
			st = append(st, num)
			num = 0
			op = string(ch)
			continue
		}
		if op == "-" {
			st = append(st, -num)
			num = 0
			op = string(ch)
			continue
		}
		if op == "*" {
			// 表达式是正确的
			last := st[len(st)-1]
			st = st[:len(st)-1]
			st = append(st, last*num)
			op = string(ch)
			num = 0
			continue
		}
		if op == "/" {
			// 表达式是正确的
			last := st[len(st)-1]
			st = st[:len(st)-1]
			st = append(st, last/num)
			op = string(ch)
			num = 0
			continue
		}
	}

	ans := 0
	for _, ch := range st {
		ans += ch
	}
	return ans
}
