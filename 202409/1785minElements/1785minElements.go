package main

import "sort"

func main() {

}

func minElements(nums []int, limit int, goal int) int {
	sort.Ints(nums)
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	d := abs(sum - goal)
	ans := (d + limit - 1) / limit
	return ans
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
