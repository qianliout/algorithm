package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isTransformable("848188", "818884"))
	fmt.Println(isTransformable("4941", "1494"))
}

// 没有明白
func isTransformable(s string, t string) bool {
	n, m := len(s), len(t)
	if n != m {
		return false
	}
	g := make([][]int, 10)
	for i, ch := range s {
		idx := int(ch) - int('0')
		g[idx] = append(g[idx], i)
	}
	for i := 0; i < n; i++ {
		dig := int(t[i]) - int('0')
		if len(g[dig]) == 0 {
			return false
		}
		for j := 0; j < dig; j++ {
			if len(g[j]) > 0 && g[j][0] < g[dig][0] {
				return false
			}
		}
		g[dig] = g[dig][1:]
	}
	return true
}
