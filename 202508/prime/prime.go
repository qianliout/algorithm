package main

import (
	"fmt"
)

func main() {
	fmt.Println(EulerPrime2(10))
	fmt.Println(Eratosthenes(10))
}

func EulerPrime2(n int) []int {

	prim := make([]int, 0)
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true // 默认全是
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prim = append(prim, i)
		}
		for _, p := range prim {
			if p*i > n {
				break
			}
			isPrime[p*i] = false
			if i%p == 0 {
				break
			}
		}
	}
	return prim
}

func Eratosthenes(n int) []int {
	prim := make([]int, 0)
	isPrim := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrim[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrim[i] {
			for j := i * i; j <= n; j += i {
				isPrim[j] = false
			}
		}
	}
	for i, b := range isPrim {
		if b {
			prim = append(prim, i)
		}
	}
	return prim
}

func Eratosthenes2(n int) []int {
	prim := make([]int, 0)
	isPrim := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrim[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrim[i] {
			for j := i * i; j <= n; j += i {
				isPrim[j] = false
			}
		}
	}
	for i, b := range isPrim {
		if b {
			prim = append(prim, i)
		}
	}
	return prim
}
