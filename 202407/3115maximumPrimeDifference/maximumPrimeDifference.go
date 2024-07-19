package main

import (
	"math"
)

func main() {

}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func maximumPrimeDifference(nums []int) int {
	le, ri := -1, -1

	for i := 0; i < len(nums); i++ {
		if isPrime(nums[i]) {
			le = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if isPrime(nums[i]) {
			ri = i
			break
		}
	}
	return ri - le
}
