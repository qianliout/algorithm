package main

import (
	"math"
)

func main() {

}

func countHomogenous(s string) int {
	mod := int(math.Pow10(9)) + 7
	i, n := 0, len(s)
	ans := 0
	for i < n {
		j := i
		for j < n && s[i] == s[j] {
			j++
		}
		cnt := j - i
		ans += (cnt + 1) * cnt / 2
		ans = ans % mod
		i = j
	}
	return ans % mod
}
