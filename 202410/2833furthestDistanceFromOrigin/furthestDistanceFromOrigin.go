package main

import (
	"strings"
)

func main() {

}

func furthestDistanceFromOrigin(moves string) int {
	l := strings.Count(moves, "L")
	r := strings.Count(moves, "R")
	n := len(moves)
	return abs(l-r) + (n - l - r)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
