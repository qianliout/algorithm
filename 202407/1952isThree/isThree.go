package main

import (
	"math"
)

func main() {

}

func isThree(n int) bool {
	cnt := 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			cnt++
			if n/i != i {
				cnt++
			}
		}
		if cnt >= 2 {
			return false
		}
	}

	return cnt == 1
}
