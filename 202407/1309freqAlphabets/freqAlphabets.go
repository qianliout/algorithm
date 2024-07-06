package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	n := len(s)
	ans := make([]string, 0)
	m := make(map[string]string)
	for i := 1; i <= 9; i++ {
		m[strconv.Itoa(i)] = string(byte(i - 1 + 'a'))
	}
	start := 10
	for i := 'j'; i <= 'z'; i++ {
		key := fmt.Sprintf("%d#", start)
		m[key] = string(i)
		start++
	}

	i := 0
	for i < n {
		if i+2 < n && s[i+2] == '#' {
			ans = append(ans, m[s[i:i+3]])
			i = i + 3
		} else {
			ans = append(ans, m[s[i:i+1]])
			i++
		}
	}
	return strings.Join(ans, "")
}
