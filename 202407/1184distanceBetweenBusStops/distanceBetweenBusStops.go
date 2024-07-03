package main

import (
	"fmt"
)

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)
	i, j := start, start
	a, b := 0, 0
	for i != destination {
		a += distance[i]
		i++
		if i == n {
			i = 0
		}
	}
	for j != destination {
		j--
		if j < 0 {
			j = n - 1
		}
		b += distance[j]
	}
	return min(a, b)
}
