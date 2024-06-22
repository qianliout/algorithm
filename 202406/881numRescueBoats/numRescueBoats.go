package main

import (
	"sort"
)

func main() {

}

/*
每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
返回 承载所有人所需的最小船数 。
贪心的做法
*/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	if limit < people[len(people)-1] {
		return 0 // 不能完成
	}
	ans, le, ri := 0, 0, len(people)-1
	for le <= ri {
		if people[le]+people[ri] <= limit {
			ans++
			le++
			ri--
		} else {
			ans++
			ri--
		}
	}
	return ans
}
