package main

func main() {

}

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
func calculate1(s string) int {
	num, op := 0, 1
	nums := make([]int, 0)
	ops := make([]int, 0)
	ans := 0
	for _, c := range s {
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
			ans += num * op
			ops = append(ops, op)
			nums = append(nums, ans)
			ans, num, op = 0, 0, 1
		case ')':
			ans += num * op
			ans = ans * ops[len(ops)-1]
			ans += nums[len(nums)-1]
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			num, op = 0, 1
		default:
			num = num*10 + int(c) - int('0')
		}
	}
	return ans + op*num
}

// 对s 求值， s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开,没有括号
func calculate(s string) int {
	nums := make([]int, 0)
	num := 0
	op := byte('+')
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		for i < len(s) {
			ch := s[i]
			if ch >= '0' && ch <= '9' {
				num = num*10 + int(ch) - 48
				i++
			} else {
				break
			}
		}
		switch op {
		case '+':
			nums = append(nums, num)
		case '-':
			nums = append(nums, -num)
		case '*':
			n := len(nums)
			last := nums[n-1]
			// fir, sec := nums[n-2], nums[n-1]
			nums = nums[:n-1]
			nums = append(nums, last*num)
		case '/':
			n := len(nums)
			last := nums[n-1]
			// fir, sec := nums[n-2], nums[n-1]
			nums = nums[:n-1]
			nums = append(nums, last/num)
		}
		// 一定要判断
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
