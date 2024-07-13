package main

import (
	"math"
)

func main() {

}

func numSub(s string) int {
	res := make([]int, 0)
	cnt := 0
	for _, ch := range s {
		if ch == '1' {
			cnt++
		} else {
			res = append(res, cnt)
			cnt = 0
		}
	}
	res = append(res, cnt)
	ans := 0
	mod := int(math.Pow10(9)) + 7
	for _, ch := range res {
		ans = (ans + ch*(ch+1)/2) % mod
	}

	return ans
}
