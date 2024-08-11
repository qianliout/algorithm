package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31"))
}

var dir = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	a := cal(max(arriveAlice, arriveBob))
	b := cal(min(leaveAlice, leaveBob))
	return max(0, b-a+1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func cal(a string) int {
	sum := make([]int, 12)
	sum[0] = dir[0]
	for i := 1; i < 12; i++ {
		sum[i] = sum[i-1] + dir[i]
	}
	split := strings.Split(a, "-")
	m, _ := strconv.Atoi(split[0])
	d, _ := strconv.Atoi(split[1])

	ans := 0
	for i := 0; i < m; i++ {
		ans += dir[i]
	}

	ans += d
	return ans
}
