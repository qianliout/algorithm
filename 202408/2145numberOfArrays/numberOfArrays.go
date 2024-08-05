package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5))

}

func numberOfArrays(differences []int, lower int, upper int) int {
	mi, mx := 0, 0 // 这里只能是赋值成0，0
	num := 0
	n := len(differences)
	for i := 0; i < n; i++ {
		num = num + differences[i]
		mi = min(mi, num)
		mx = max(mx, num)
	}
	return max(0, upper-lower+1-(mx-mi))
}
