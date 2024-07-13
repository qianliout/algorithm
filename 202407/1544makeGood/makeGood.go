package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeGood("abBAcC"))
}

func makeGood(s string) string {
	ss := []byte(s)
	base := abs(int('a') - int('A'))
	ans := make([]byte, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		for i < n && len(ans) > 0 && abs(int(ss[i])-int(ans[len(ans)-1])) == base {
			ans = ans[:len(ans)-1]
			i++
		}
		if i < n {
			ans = append(ans, ss[i])
		}
	}
	return string(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
