package main

import (
	"strconv"
)

func main() {

}

func maximumValue(strs []string) int {
	mx := 0
	for _, ch := range strs {
		mx = max(mx, cal(ch))
	}
	return mx
}

func check(s string) bool {
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && '9' >= s[i] {
			continue
		}
		return false
	}
	return true
}

func cal(s string) int {
	if check(s) {
		i, _ := strconv.Atoi(s)
		return i
	}
	return len(s)
}
