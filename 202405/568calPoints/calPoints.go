package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}

func calPoints(operations []string) int {

	/*
		整数 x - 表示本回合新获得分数 x
		"+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
		"D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
		"C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数

	*/
	stark := make([]int, 0)
	for _, ch := range operations {
		switch ch {
		case "+":
			fir, sec := stark[len(stark)-1], stark[len(stark)-2]
			stark = append(stark, fir+sec)
		case "D":
			last := stark[len(stark)-1]
			stark = append(stark, last*2)
		case "C":
			stark = stark[:len(stark)-1]
		default:
			n, _ := strconv.Atoi(ch)
			stark = append(stark, n)
		}
	}
	ans := 0
	for _, ch := range stark {
		ans += ch
	}
	return ans
}
