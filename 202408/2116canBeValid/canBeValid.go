package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeValid("))()))", "010100"))
}

func canBeValid2(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' || locked[i] == '0' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			return false
		}
	}
	cnt = 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			return false
		}
	}
	return true
}

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' || locked[i] == '0' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			return false
		}
	}
	cnt = 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			return false
		}
	}
	return true
}
