package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestRepunitDivByK(23))
}

func smallestRepunitDivByK(k int) int {
	seen := make(map[int]int)
	i := 1 % k
	for i > 0 && seen[i] {
		if i%k == 0 {
			return len(seen) + 1
		}
		i = i*10 + 1
		seen[i%k]++
		if len(seen) >= k-1 {
			return -1
		}
		// if seen[i%k] > 0 {
		// 	break
		// }

	}
	// return -1
}
