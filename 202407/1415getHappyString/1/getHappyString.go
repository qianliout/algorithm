package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getHappyString(3, 7))
	fmt.Println(getHappyString(3, 9))
	fmt.Println(getHappyString(2, 7))
}

func getHappyString(n int, k int) string {
	if cal(n)*3 < k {
		return ""
	}
	stateTab := [][]int{{1, 2}, {0, 2}, {0, 1}}
	ans := make([]byte, n)
	order := k - 1
	idx := 0
	state := order / cal(n)
	ans[idx] = byte(state + 'a')
	idx++

	tree := order & (1<<(n-1) - 1)
	for i := n - 2; i >= 0; i-- {
		state = stateTab[state][(tree>>i)&1]
		ans[idx] = byte(state + 'a')
		idx++
	}

	return string(ans)
}

func cal(n int) int {
	return int(math.Pow(2, float64(n-1)))
}

func redu(apl []byte, idx int) []byte {
	ans := []byte{}
	for i, b := range apl {
		if idx == i {
			continue
		}
		ans = append(ans, b)
	}
	return ans

}
