package main

import (
	"fmt"
)

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	d := make([]int, n+1)
	for _, ch := range bookings {
		x, y, z := ch[0]-1, ch[1]-1, ch[2]
		d[x] += z
		// 这里是重点，一定y+1
		d[y+1] -= z
	}

	ans := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += d[i]
		ans[i] = sum
	}

	return ans
}
