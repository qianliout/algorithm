package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minimizeResult("247+38"))
}

func minimizeResult(expression string) string {
	n := len(expression)
	idx := strings.Index(expression, "+")
	mx := math.MaxInt
	ans := ""
	for i := 0; i < idx; i++ {
		fir := expression[:i]
		for j := idx + 1; j < n; j++ {
			sec := expression[i : j+1]
			thid := expression[j+1:]
			c := cal2(fir, sec, thid)
			if c < mx {
				mx = c
				ans = fmt.Sprintf("%s(%s)%s", fir, sec, thid)
			}
		}
	}
	return ans
}

func cal(s string) int {
	idx := strings.Index(s, "+")
	if idx == -1 {
		atoi, _ := strconv.Atoi(s)
		return atoi
	}
	a, _ := strconv.Atoi(s[:idx])
	b, _ := strconv.Atoi(s[idx+1:])
	return a + b
}

func cal2(s1, s2, s3 string) int {
	a, b, c := cal(s1), cal(s2), cal(s3)
	ans := b
	if s1 != "" {
		ans = ans * a
	}
	if s3 != "" {
		ans = ans * c
	}
	return ans
}
