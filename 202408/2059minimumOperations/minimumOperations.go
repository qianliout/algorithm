package main

import (
	"fmt"
)

func main() {
	//fmt.Println(minimumOperations([]int{2, 4, 12}, 2, 12))
	//fmt.Println(minimumOperations([]int{3, 5, 7}, 0, -4))
	//fmt.Println(minimumOperations([]int{2, 8, 16}, 0, 1))
	fmt.Println(minimumOperations([]int{-574938140, 347713845, 925500837, -396559946, -39213216, -696511059, -701862040,
		-547815957, -613314611, 814380075, 446824702, 397447568, 709912715, 144793599, 812441247, -59489753, 857299470,
		360355629, 85411951, -439873837, -477453514, -887964831, -914549223, 633449658, 452658511, 657134722, 1}, 827, -547815957))
}

func minimumOperations(nums []int, start int, goal int) int {
	if start == goal {
		return 0
	}
	ans := 0
	queue := []int{start}
	cnt := make(map[int]bool)
	cnt[start] = true
	for len(queue) > 0 {
		lev := make([]int, 0)
		ans++
		for _, qu := range queue {
			for _, ch := range nums {
				a := qu + ch
				b := qu - ch
				c := qu ^ ch

				if a == goal || b == goal || c == goal {
					return ans
				}
				if a >= 0 && a <= 1000 && !cnt[a] {
					cnt[a] = true
					lev = append(lev, a)
				}
				if b >= 0 && b <= 1000 && !cnt[b] {
					cnt[b] = true
					lev = append(lev, b)
				}
				if c >= 0 && c <= 1000 && !cnt[c] {
					cnt[c] = true
					lev = append(lev, c)
				}
			}
		}
		queue = lev
	}

	return -1
}

/*
如果 0 <= x <= 1000 ，那么，对于数组中的任一下标 i（0 <= i < nums.length），可以将 x 设为下述任一值：
    x + nums[i]
    x - nums[i]
    x ^ nums[i]（按位异或 XOR）
注意，你可以按任意顺序使用每个 nums[i] 任意次。使 x 越过 0 <= x <= 1000 范围的运算同样可以生效，但该该运算执行后将不能执行其他运算。
返回将 x = start 转化为 goal 的最小操作数；如果无法完成转化，则返回 -1 。
*/
