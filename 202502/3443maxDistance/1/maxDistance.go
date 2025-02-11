package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistance("NWSE", 1))
	fmt.Println(maxDistance("NSWWEW", 3))
}

// 不能得到正确的结果，不知道为啥
func maxDistance(s string, k int) int {
	n := len(s)
	ss := []byte(s)
	var dfs func(i, u int) []position
	dfs = func(i, u int) []position {
		if i == 0 {
			ans := []position{add(position{}, ss[i])}
			if u > 0 {
				for _, c := range not(ss[i]) {
					b := add(position{}, c)
					ans = append(ans, b)
				}
			}
			return ans
		}

		// not change
		res := make([]position, 0)
		pre := dfs(i-1, u)
		for _, c := range pre {
			res = append(res, add(c, ss[i]))
		}

		// change
		if u > 0 {
			bb := dfs(i-1, u-1)
			for _, b := range bb {
				for _, c := range not(ss[i]) {
					res = append(res, add(b, c))
				}
			}
		}
		return res
		// return getMax(res)
	}

	ans := dfs(n-1, k)
	mx := getMax(ans)
	if len(mx) == 0 {
		return 0
	}
	return cal(mx[0])
}

type position struct {
	x, y int
}

func getMax(ss []position) []position {
	mx := 0
	ans := make([]position, 0)
	for _, s := range ss {
		c := cal(s)
		if c < mx {
			continue
		}
		if c == mx {
			ans = append(ans, s)
			continue
		}
		if c > mx {
			mx = c
			ans = []position{s}
		}
	}
	return ans
}

func not(a byte) []byte {
	ss := []byte("NSEW")
	ans := make([]byte, 0)
	for _, c := range ss {
		if c != a {
			ans = append(ans, c)
		}
	}
	return ans
}

func add(p position, c byte) position {
	ans := position{
		x: p.x,
		y: p.y,
	}
	switch c {
	case 'N':
		ans.x++
	case 'S':
		ans.x--
	case 'E':
		ans.y--
	case 'W':
		ans.y++
	}
	return ans
}

func cal(p position) int {
	return abs(p.x) + abs(p.y)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
