package main

import (
	"fmt"
)

func main() {
	fmt.Println(clumsy(10))
}

func clumsy(n int) int {
	stark := make([]int, n+1)
	stark = append(stark, n)
	op := 0
	for i := n - 1; i >= 1; i-- {
		if op == 0 {
			last := stark[len(stark)-1]
			stark[len(stark)-1] = last * i
			op = 1
		} else if op == 1 {
			last := stark[len(stark)-1]
			stark[len(stark)-1] = last / i
			op = 2
		} else if op == 2 {
			stark = append(stark, i)
			op = 3
		} else if op == 3 {
			stark = append(stark, -i)
			op = 0
		}
	}
	sum := 0
	for _, ch := range stark {
		sum += ch
	}
	return sum
}
