package main

import (
	"fmt"
)

func main() {
	fmt.Println(replaceNonCoprimes([]int{6, 4, 3, 2, 7, 6, 2})) // 12,7,6
}

func replaceNonCoprimes(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		ans = append(ans, num)
		for len(ans) >= 2 {
			x, y := ans[len(ans)-2], ans[len(ans)-1]
			if gcd(x, y) == 1 {
				break
			} else {
				ans = ans[:len(ans)-1]
				ans[len(ans)-1] = lcm(x, y)
			}
		}
	}
	return ans
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
