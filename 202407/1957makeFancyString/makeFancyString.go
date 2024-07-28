package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeFancyString("leeetcode"))
}

func makeFancyString(s string) string {
	i := 0
	ss := []byte(s)
	n := len(s)
	ans := make([]byte, 0)
	for i < n {
		j := i + 1
		for j < n && s[j] == s[i] {
			j++
		}
		m := min(n, i+min(2, j-i))
		ans = append(ans, ss[i:m]...)
		i = j
	}
	return string(ans)
}
