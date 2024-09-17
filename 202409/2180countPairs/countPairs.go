package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 2))
}

func countPairs(nums []int, k int) int64 {
	d := make([]int, 0)
	for i := 1; i*i <= k; i++ {
		if k%i == 0 {
			d = append(d, i)
			if i*i < k {
				d = append(d, k/i)
			}
		}
	}
	cnt := make(map[int]int)
	ans := 0
	for _, num := range nums {
		ans += cnt[k/gcd(num, k)]
		for _, i := range d {
			if num%i == 0 {
				cnt[i]++
			}
		}
	}
	return int64(ans)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

// func gcd(a, b int) int {
// 	for a != 0 {
// 		a, b = b%a, a
// 	}
// 	return b
// }
