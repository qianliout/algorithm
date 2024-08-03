package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
}

func addSpaces(s string, spaces []int) string {
	ss := []byte(s)
	ans := make([]string, 0)
	n := len(spaces)
	for i := 0; i < len(spaces); i++ {
		pre := 0
		if i > 0 {
			pre = spaces[i-1]
		}
		now := spaces[i]
		ans = append(ans, string(ss[pre:now]))
	}

	if spaces[n-1] < len(s) {
		ans = append(ans, string(ss[spaces[n-1]:]))
	}

	return strings.Join(ans, " ")
}
