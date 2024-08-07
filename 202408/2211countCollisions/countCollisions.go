package main

import (
	"strings"
)

func main() {

}

func countCollisions(directions string) int {
	directions = strings.TrimLeft(directions, "L")
	directions = strings.TrimRight(directions, "R")
	return len(directions) - strings.Count(directions, "S")
}
