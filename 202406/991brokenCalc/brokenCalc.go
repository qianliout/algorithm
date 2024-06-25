package main

import (
	"fmt"
)

func main() {
	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))
}

func brokenCalc(startValue int, target int) int {
	if startValue > target {
		return startValue - target
	}
	if startValue < target {
		if target%2 == 0 {
			return brokenCalc(startValue, target/2) + 1
		} else {
			return brokenCalc(startValue, target+1) + 1
		}
	}
	return 0
}
