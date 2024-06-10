package main

import (
	"fmt"
)

func main() {
	fmt.Println(getAverages([]int{8}, 1000))
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
}

func getAverages(nums []int, k int) []int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	// 也可以不判断
	// if n < 2*k+1 {
	// 	return ans
	// }

	for i := k; i+k < n; i++ {
		sub := sum[i+k+1] - sum[i-k]
		ans[i] = sub / ((k)*2 + 1)
	}

	return ans
}

func minimumRecolors(blocks string, k int) int {
	ss := []byte(blocks)
	n, b := len(ss), 0
	ans := n
	le, ri := 0, 0
	for le <= ri && ri < n {
		if ss[ri] == 'B' {
			b++
		}
		ri++
		for ri-le > k {
			if ss[le] == 'B' {
				b--
			}
			le++
		}
		if ri-le == k {
			ans = min(ans, k-b)
		}
	}
	return ans
}
