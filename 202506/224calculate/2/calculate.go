package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate1("1+1"))
}

// 对s 求值， s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开,没有括号
func calculate2(s string) int {
	num := 0
	sign := byte('+')
	nums := make([]int, 0)
	n := len(s)
	// 不能写成 for i :=range s
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		for i < n && s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]) - int('0')
			i++
		}
		switch sign {
		case '+':
			nums = append(nums, num)
		case '-':
			nums = append(nums, -num)
		case '*':
			l := len(nums)
			nums[l-1] = nums[l-1] * num
		case '/':
			l := len(nums)
			nums[l-1] = nums[l-1] / num
		}
		if i < n {
			sign, num = byte(s[i]), 0
		}
	}
	ans := 0
	for _, c := range nums {
		ans += c
	}
	return ans
}

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
func calculate1(s string) int {
	op, num := 1, 0
	nums := make([]int, 0)
	ops := make([]int, 0)
	ans := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case ' ':
			continue
		case '+':
			ans += op * num
			op, num = 1, 0
		case '-':
			ans += op * num
			op, num = -1, 0
		case '(':
			ans += op * num
			nums = append(nums, ans)
			ops = append(ops, op)
			op, num, ans = 1, 0, 0
		case ')':
			ans += op * num
			ans *= ops[len(ops)-1]
			ans += nums[len(nums)-1]
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			op, num = 1, 0
		default:
			num = num*10 + int(c) - int('0')
		}
	}
	return ans + op*num
}
