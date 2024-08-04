package main

import (
	"sort"
)

func main() {

}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, ch := range asteroids {
		if mass < ch {
			return false
		}
		mass += ch
	}

	return true
}
