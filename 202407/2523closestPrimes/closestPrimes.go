package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(closestPrimes(10, 19))
	fmt.Println(closestPrimes(1, 1000000))
}

var primes []int
var mx int = int(math.Pow10(6)) + 1

// 欧筛法
// func init() {
// 	all := make([]bool, mx+10)
// 	for i := 2; i <= mx; i++ {
// 		if !all[i] {
// 			primes = append(primes, i)
// 		}
// 		for j := i * i; j <= mx; j = j + i {
// 			all[j] = true
// 		}
// 	}
// 	// primes = append(primes, mx)
// 	// primes = append(primes, mx)
// }

// 线筛选法，也叫欧拉筛选
func init() {
	all := make([]bool, mx+1)
	for i := 2; i <= mx; i++ {
		if !all[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > mx {
				break
			}
			all[p*i] = true
			if i%p == 0 {
				break
			}
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func closestPrimes(left int, right int) []int {
	le, ri := left-1, right+1

	for i := left; i <= right; i++ {
		if isPrime(i) {
			le = i
			break
		}
	}
	for i := right; i >= left; i-- {
		if isPrime(i) {
			ri = i
			break
		}
	}
	return []int{le, ri}
}
