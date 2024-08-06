package main

import (
	"fmt"
)

func main() {
	fmt.Println(countOperations(2, 3))
}

func countOperations(num1 int, num2 int) int {
	ans := 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			a := num1 - num2
			num1, num2 = max(a, num2), min(a, num2)
		} else {
			a := num2 - num1
			num1, num2 = max(a, num1), min(a, num1)
		}
		ans++
	}

	return ans
}
