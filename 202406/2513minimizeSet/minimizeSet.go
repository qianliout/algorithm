package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimizeSet(2, 7, 1, 3))
	fmt.Println(minimizeSet(3, 5, 2, 1))
}

/*
能被 d1整除，不能被d2整除的数 只能在 arr2 中，不能在 arr1中，这部分数记做 A
能被 d1整除，不能被d1整除的数 只能在 arr1 中,不能在 arr2中 这部分数记录 B
既能被 d1整除，又能被 d2整除的数（也就是能被最大公约数整除的数），arr1,arr2都不能存在) 这部分数记做 C

剩下的数就是都可以
*/
func minimizeSet(d1 int, d2 int, u1 int, u2 int) int {
	lcm := CalLcm(d1, d2)
	var check func(x int) bool
	check = func(x int) bool {
		// 先用只能填 arr1的数填充arr1,看还剩下多少数没有填充
		left1 := max(0, u1-(x/d2-x/lcm))
		// 先用只能填 arr1的数填充arr1,看还剩下多少数没有填充
		left2 := max(0, u2-(x/d1-x/lcm))
		// 这部分就是都可以用的
		common := x - x/d1 - x/d2 + x/lcm
		// 如果都可以用的数可以填充满，就说明 x 这个数满足要求
		return common >= left1+left2
	}
	mi, mx := 0, math.MaxInt/10
	le, ri := 0, mx

	for le < ri {
		// 左端点
		mid := le + (ri-le)/2
		if mid >= mi && mid < mx && check(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

// 求最大公约数
func CalLcm(a, b int) int {
	x, y := a, b
	for a != 0 {
		a, b = b%a, a
	}
	return x * y / b
}
