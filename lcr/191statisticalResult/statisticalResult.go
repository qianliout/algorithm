package main

import (
	"fmt"
)

func main() {
	fmt.Println(statisticalResult([]int{1, 2, 0, 4, 5}))
}

func statisticalResult(arrayA []int) []int {
	n := len(arrayA)
	ans := make([]int, n)
	if n <= 1 {
		return ans
	}
	zeroCnt := 0
	for i := 0; i < n; i++ {
		if arrayA[i] == 0 {
			zeroCnt++
		}
	}
	if zeroCnt > 1 {
		return ans
	}

	all := 1
	for _, v := range arrayA {
		if v == 0 {
			continue
		}
		all *= v
	}
	for i := 0; i < n; i++ {
		if arrayA[i] == 0 {
			ans[i] = all
			continue
		}
		if zeroCnt > 0 {
			ans[i] = 0
		} else {
			ans[i] = all / arrayA[i]
		}
	}

	return ans
}
