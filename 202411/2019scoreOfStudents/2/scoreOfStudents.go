package main

import (
	"fmt"
)

func main() {
	fmt.Println(cal("7+3*1*2"))
	fmt.Println(cal2("7+3*1*2"))
	fmt.Println(cal2("3*3*2+2*5+0*0+1*6+0*4+4*0+2*2"))
}

func scoreOfStudents(s string, answers []int) int {
	ok := cal(s)
	notOK := cal2(s)
	cnt := make(map[int]bool)
	for _, ch := range notOK {
		if ch != ok {
			cnt[ch] = true
		}
	}
	ans := 0
	for _, ch := range answers {
		if ch == ok {
			ans += 5
		} else if cnt[ch] {
			ans += 2
		}
	}

	return ans
}

func cal(s string) int {
	if s == "" {
		return 0
	}
	st := []int{0}
	op := '+'
	for _, ch := range s {
		if ch == '+' {
			op = '+'
			continue
		}
		if ch == '*' {
			op = '*'
			continue
		}
		// 表达式中所有整数运算数字都在闭区间 [0, 9] 以内。
		// 所以可以直接算
		num := int(ch) - int('0')
		if op == '+' {
			st = append(st, num)
		} else if op == '*' {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			st = append(st, last*num)
			op = '+'
		}
	}
	ans := 0
	for _, ch := range st {
		ans += ch
	}
	return ans
}

func cal2(s string) []int {
	if s == "" {
		return []int{}
	}
	if len(s) == 1 && s[0] >= '0' && s[0] <= '9' {
		return []int{int(s[0]) - int('0')}
	}

	ans := make([]int, 0)
	for i, ch := range s {
		if ch == '+' {
			ans = append(ans, cal(s[:i])+cal(s[i+1:]))
		}
		if ch == '*' {
			ans = append(ans, cal(s[:i])*cal(s[i+1:]))
		}
	}
	res := make([]int, 0)
	cnt := make(map[int]bool)
	for _, ch := range ans {
		if !cnt[ch] {
			res = append(res, ch)
			cnt[ch] = true
		}
	}

	return res
}
