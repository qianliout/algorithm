package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(crackPassword([]int{0, 3, 30, 34, 5, 9}))
}

func crackPassword(password []int) string {
	n := len(password)
	pp := make([]string, n)
	for i, ch := range password {
		pp[i] = fmt.Sprintf("%d", ch)
	}
	sort.Slice(pp, func(i, j int) bool {
		a := fmt.Sprintf("%s%s", pp[i], pp[j])
		b := fmt.Sprintf("%s%s", pp[j], pp[i])
		return a < b
	})
	ans := strings.Join(pp, "")
	return ans
}

// 关键在于排序规则：不是直接比较两个数字的大小，而是比较两个数字拼接后的字符串哪一个更小。例如，对于数字 3 和 30，我们应该比较 "303" 和 "330" 来决定谁应该排在前面。
