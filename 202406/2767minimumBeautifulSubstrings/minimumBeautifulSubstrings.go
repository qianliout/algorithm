package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumBeautifulSubstrings("1011"))
	fmt.Println(minimumBeautifulSubstrings("111"))
	fmt.Println(minimumBeautifulSubstrings("0"))
}

func minimumBeautifulSubstrings(s string) int {
	inf := math.MaxInt / 2
	n := len(s)
	ss := []byte(s)
	var dfs func(i int) int
	pow5 := gen()
	// 以i开头的字符串可以分成多少
	dfs = func(i int) int {
		if i == n || i < 0 {
			return 0
		}
		res := inf
		if ss[i] == '0' {
			// 这里一定不能返回-1
			return res
		}

		for _, t := range pow5 {
			if i+len(t) > n {
				break
			}
			if string(ss[i:i+len(t)]) == t {
				res = min(res, dfs(i+len(t))+1)
			}
		}
		return res
	}
	a := dfs(0)
	if a >= inf {
		return -1
	}
	return a
}

func gen() []string {
	ans := make([]string, 0)
	i := 0
	for {
		a := int(math.Pow(5, float64(i)))
		if a > 1<<18 {
			break
		}
		ans = append(ans, fmt.Sprintf("%b", a))
		i++
	}
	return ans
}
