package main

import (
	"fmt"
)

func main() {
	fmt.Println(sockCollocation([]int{1, 2, 4, 1, 4, 3, 12, 3}))
	fmt.Println(sockCollocation([]int{1, 2, 5, 2}))
}

func sockCollocation(sockets []int) []int {
	fl := 0
	for _, ch := range sockets {
		fl = fl ^ ch
	}
	// 只保留最后一位1
	fl = fl & -fl

	left, right := 0, 0
	for _, ch := range sockets {
		// 剩下的两个数，在这一位上肯定是一个有1，有个是0
		if ch&fl > 0 {
			left = left ^ ch
		} else {
			right = right ^ ch
		}
	}
	return []int{left, right}
}

func trainingPlan(actions []int) int {
	num := 0
	for i := 0; i < 64; i++ {
		cnt := 0
		for _, ch := range actions {
			if (ch>>i)&1 > 0 {
				cnt++
			}
		}
		if cnt%3 != 0 {
			// 两种写法都可以
			num = num + (1 << i)
			// num = num ^ (1 << i)
		}
	}
	return num
}
