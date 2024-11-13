package main

func main() {

}

func calculate(s string) int {
	nums := make([]int, 0)

	ans, num, op := 0, 0, 1
	ops := make([]int, 0)
	for _, ch := range s {
		switch ch {
		case ' ':
			continue
		case '+':
			ans += num * op
			num, op = 0, 1
		case '-':
			ans += num * op
			num, op = 0, -1
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
			num = num*10 + int(ch-'0')
		}
	}
	return ans + op*num
}

// 1 <= s.length <= 3 * 105
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
// s 表示一个有效的表达式
// '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
// '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
// 输入中不存在两个连续的操作符
// 每个数字和运行的计算将适合于一个有符号的 32位 整数
