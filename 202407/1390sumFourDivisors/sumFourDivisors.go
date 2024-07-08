package main

import (
	"math"
)

func main() {

}

func sumFourDivisors(nums []int) int {
	ans := 0
	for _, ch := range nums {
		res := divisor(ch)
		if len(res) == 4 {
			ans += Sum(res)
		}
	}
	return ans
}

func Sum(nums []int) int {
	ans := 0
	for _, ch := range nums {
		ans += ch
	}
	return ans

}

func divisor(num int) []int {
	ans := make([]int, 0)
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			ans = append(ans, i)
			if num/i != i {
				ans = append(ans, num/i)
			}
		}
	}
	return ans
}
