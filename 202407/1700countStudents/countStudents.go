package main

import (
	"fmt"
)

func main() {
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}

func countStudents(students []int, sandwiches []int) int {
	cnt := make([]int, 2)
	for _, ch := range students {
		cnt[ch]++
	}
	for i := 0; i < len(sandwiches); i++ {
		sa := sandwiches[i]
		cnt[sa]--
		if cnt[sa] == -1 {
			return len(sandwiches) - i
		}
	}
	return 0
}

// 根据题意进行模拟即可 : 当学生遇到喜欢的种类会进行匹配，否则会轮到队列尾部，而面包则是一直停在栈顶位置等待匹配。
// 因此当且仅当栈顶的面包种类没有待匹配的学生种类与之相对应时，整个匹配过程结束
