package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(getHappyString(3, 7))
	// fmt.Println(getHappyString(3, 9))
	// fmt.Println(getHappyString(2, 7))
	fmt.Println(getHappyString(3, 1))
}

func getHappyString(n int, k int) string {
	leng := n
	if cal(n)*3 < k {
		return ""
	}
	k = k - 1
	apl := []byte{'a', 'b', 'c'}
	ans := make([]byte, 0)
	for k >= 0 {
		m := cal(n)
		idx := k / m
		ans = append(ans, apl[idx])
		if len(ans) >= leng {
			break
		}
		k -= m * (idx)
		n--
		apl = redu(apl, idx)
	}

	return string(ans)
}

func cal(n int) int {
	return int(math.Pow(2, float64(n-1)))
}

func redu(pre []byte, idx int) []byte {
	red := pre[idx]
	apl := []byte{'a', 'b', 'c'}

	ans := []byte{}
	for _, b := range apl {
		if b == red {
			continue
		}
		ans = append(ans, b)
	}
	return ans
}
