package main

import (
	"fmt"
)

func main() {

}

func simplifiedFractions(n int) []string {
	ans := make([]string, 0)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcb(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return ans
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}
