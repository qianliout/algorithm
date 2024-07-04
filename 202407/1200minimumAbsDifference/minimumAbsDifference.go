package main

import (
	"math"
	"sort"
)

func main() {

}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	sub := math.MaxInt / 10
	res := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		a := abs(arr[i] - arr[i-1])
		if a < sub {
			sub = a
			res = make([][]int, 0)
			res = append(res, []int{arr[i-1], arr[i]})
		} else if a > sub {
			continue
		} else if a == sub {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][1] })
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
