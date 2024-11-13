package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func evalRPN(tokens []string) int {
	nums := make([]int, 0)
	for _, ch := range tokens {
		if !operator(ch) {
			nums = append(nums, getNum(ch))
			continue
		}

		n := len(nums)
		fir := nums[n-2]
		sec := nums[n-1]
		nums = nums[:n-2]
		switch ch {
		case "+":
			nums = append(nums, fir+sec)
		case "-":
			nums = append(nums, fir-sec)
		case "*":
			nums = append(nums, fir*sec)
		case "/":
			nums = append(nums, fir/sec)
		}
	}
	ans := 0
	for _, ch := range nums {
		ans += ch
	}
	return ans
}

func operator(ch string) bool {
	if ch == "+" || ch == "-" || ch == "*" || ch == "/" {
		return true
	}
	return false
}

func getNum(ch string) int {
	i, _ := strconv.Atoi(ch)
	return i
}
