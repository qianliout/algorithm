package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSteps("1101"))
}

func numSteps(s string) int {
	cnt := 0
	for !(len(s) == 1 && s[0] == '1') {
		if s[len(s)-1] == '0' {
			s = s[:len(s)-1]
		} else {
			s = add(s)
		}
		cnt++
	}

	return cnt
}

func add(s string) string {
	a := 1
	n := len(s)
	ans := make([]byte, n+1)
	for i := n - 1; i >= 0; i-- {
		b := a + int(s[i]-'0')
		if b == 1 {
			ans[i+1] = '1'
			a = 0
		} else if b == 0 {
			ans[i+1] = '0'
			a = 0
		} else {
			ans[i+1] = '0'
			a = 1
		}
	}
	if a == 1 {
		ans[0] = '1'
	} else {
		ans = ans[1:]
	}
	return string(ans)
}
