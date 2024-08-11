package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countTime("??:??"))
	fmt.Println(countTime("?5:00"))
}

func countTime(ti string) int {
	split := strings.Split(ti, ":")
	a, b := 0, 0

	for i := 0; i < 24; i++ {
		h := strconv.Itoa(i)
		if len(h) == 1 {
			h = "0" + h
		}
		if check(h, split[0]) {
			a++
		}
	}
	for i := 0; i < 60; i++ {
		h := strconv.Itoa(i)
		if len(h) == 1 {
			h = "0" + h
		}
		if check(h, split[1]) {
			b++
		}
	}
	return a * b
}

func check(a, b string) bool {
	for i := 0; i < 2; i++ {
		if b[i] == '?' {
			continue
		}
		if a[i] == b[i] {
			continue
		}
		return false

	}
	return true
}
