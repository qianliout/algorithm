package main

import (
	"sort"
)

func main() {

}

func kthFactor(n int, k int) int {
	fac := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fac = append(fac, i)
			if n/i != i {
				fac = append(fac, n/i)
			}
		}
	}
	if len(fac) < k {
		return -1
	}
	sort.Ints(fac)
	return fac[k-1]
}
