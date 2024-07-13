package main

import (
	"fmt"
)

func main() {
	fmt.Println(numWaterBottles(9, 3))
	fmt.Println(numWaterBottles(15, 4))
}

func numWaterBottles(numBottles int, numExchange int) int {
	a := numBottles
	for numBottles >= numExchange {
		b := numBottles / numExchange
		a += b
		numBottles = numBottles%numExchange + b
	}
	return a
}
