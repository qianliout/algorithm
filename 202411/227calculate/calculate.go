package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate("3+2*2"))
}

func calculate(s string) int {
	nums := make([]int, 0)
	num := 0
	op := byte('+')
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			num = num*10 + int(s[i]) - int('0')
		}
		switch op {
		case '+':
			nums = append(nums, num)
		case '-':
			nums = append(nums, -num)
		case '*':
			n := len(nums)
			last := nums[n-1]
			nums = nums[:n-1]
			nums = append(nums, last*num)
		case '/':
			n := len(nums)
			last := nums[n-1]
			nums = nums[:n-1]
			nums = append(nums, last/num)
		}
		// 这是容易出错点
		if i < len(s) {
			num, op = 0, s[i]
		}
	}

	ans := 0
	for _, ch := range nums {
		ans += ch
	}
	return ans
}

// 1 <= s.length <= 3 * 105
// s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
// s 表示一个 有效表达式
// 表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
// 题目数据保证答案是一个 32-bit 整数
