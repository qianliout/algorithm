package main

import (
	"fmt"
)

func main() {
	fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
	fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2}))

}
func findEvenNumbers(digits []int) []int {
	cnt := make([]int, 10)
	for _, ch := range digits {
		cnt[ch]++
	}
	ans := make([]int, 0)
	for i := 100; i <= 999; i = i + 2 {
		if check(i, cnt) {
			ans = append(ans, i)
		}
	}

	return ans
}

func check(num int, cnt []int) bool {
	mp := make([]int, 10)
	for num > 0 {
		mp[num%10]++
		num /= 10
	}
	for i := 0; i < 10; i++ {
		if mp[i] > cnt[i] {
			return false
		}
	}

	return true
}
