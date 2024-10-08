package main

import "math"

func main() {

}

func valueAfterKSeconds(n int, k int) int {
	mod := int(math.Pow10(9)) + 7
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	for k > 0 {
		for i := 1; i < n; i++ {
			ans[i] = (ans[i] + ans[i-1]) % mod
		}
		k--
	}

	return ans[n-1]
}
