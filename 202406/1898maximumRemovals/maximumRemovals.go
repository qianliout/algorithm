package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumRemovals("qlevcvgzfpryiqlwy", "qlecfqlw", []int{12, 5}))
}

// 会超时
func maximumRemovals1(s string, p string, removable []int) int {
	pp := []byte(p)
	var check func(mx int) bool
	mem := make(map[int]bool)

	ss := []byte(s)
	check = func(mx int) bool {
		for t := 0; t < mx; t++ {
			if va, ok := mem[t]; ok {
				if va {
					continue
				} else {
					return false
				}
			}
			i := removable[t]
			ss[i] = '0'
			ans := sub(ss, pp)
			mem[t] = ans
			if !ans {
				return false
			}
		}
		return true
	}
	// 开始二分了
	n := len(removable)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le+1)/2
		if le >= 0 && le < n && check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	if le > n || le < 0 || !check(le) {
		return -1
	}
	return le
}

func maximumRemovals(s string, p string, removable []int) int {
	ns, np, n := len(s), len(p), len(removable)
	var check func(mx int) bool
	check = func(mx int) bool {
		state := make([]bool, ns)
		for i := 0; i < mx; i++ {
			state[removable[i]] = true // 移除了
		}
		i, j := 0, 0
		for i < ns && j < np {
			if !state[i] && s[i] == p[j] {
				j++
			}
			i++
		}
		return j == np
	}
	// 开始二分
	le, ri := 0, n+1
	for le < ri {
		mid := le + (ri-le+1)/2
		if mid >= 0 && mid < n+1 && check(mid) {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	if le > n || le < 0 || !check(le) {
		return -1
	}
	return le
}

func sub(s []byte, p []byte) bool {
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if s[i] < 'a' {
			i++
			continue
		}
		if s[i] == p[j] {
			i++
			j++
			continue
		}
		i++
	}
	return j == len(p)
}
