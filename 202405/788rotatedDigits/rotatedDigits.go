package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(rotatedDigits(857))
}

/*
根据题目的要求，一个数是好数，当且仅当：

	数中没有出现 3,4,7；
	数中至少出现一次 2 或 5 或 6 或 9；
	对于 0,1,8 则没有要求。
*/
func rotatedDigits(n int) int {
	// 0125689
	// 2569
	ans := 0
	for i := 1; i <= n; i++ {
		if check(i) {
			ans++
		}
	}

	return ans
}

var dirs = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func check(i int) bool {
	valid, diff := true, false
	// a := map[int]bool{3: true, 4: true, 7: true}
	// b := map[int]bool{2: true, 5: true, 6: true, 9: true}
	ch := strconv.Itoa(i)
	for _, by := range ch {
		if dirs[by-'0'] == -1 {
			valid = false
		}

		if dirs[int(by)-'0'] == 1 {
			diff = true
		}
	}

	return valid && diff
}
