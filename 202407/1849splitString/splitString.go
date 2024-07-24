package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(splitString("1234"))
	fmt.Println(splitString("050043"))
	fmt.Println(splitString("10009998"))
}

func splitString(s string) bool {
	ss := []byte(s)
	n := len(s)
	var dfs func(i int, pre int) bool
	dfs = func(i int, pre int) bool {
		if i >= n {
			return true
		}
		for j := i; j < n; j++ {
			ch, _ := strconv.Atoi(string(ss[i : j+1]))
			if pre-ch == 1 {
				nex := dfs(j+1, ch)
				if nex {
					return true
				}
			} else if pre <= ch {
				break
			}
		}
		return false
	}

	// 至少要有两个数，所以 i不能到 n
	for i := 1; i < n; i++ {
		ch, _ := strconv.Atoi(string(ss[:i]))
		if dfs(i, ch) {
			return true
		}
	}
	return false
}
