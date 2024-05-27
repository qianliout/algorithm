package main

import (
	"fmt"
)

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

/*
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false
*/
func validateStackSequences(pushed []int, popped []int) bool {
	stark := make([]int, 0)
	j := 0
	for _, v := range pushed {
		stark = append(stark, v)
		for len(stark) > 0 && j < len(popped) && stark[len(stark)-1] == popped[j] {
			stark = stark[:len(stark)-1]
			j++
		}
	}
	return len(stark) == 0
}
