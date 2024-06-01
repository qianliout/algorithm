package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numFriendRequests([]int{16, 17, 18}))
}

/*
如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：
ages[y] <= 0.5 * ages[x] + 7
ages[y] > ages[x]
ages[y] > 100 && ages[x] < 100
*/
// 直接模拟会超时
func numFriendRequests2(ages []int) int {
	ans := 0
	for i, x := range ages {
		for j, y := range ages {
			if i == j {
				continue
			}
			if y <= (x/2)+7 {
				continue
			}
			if y > x {
				continue
			}
			if y > 100 && x < 100 {
				continue
			}
			ans++
		}
	}
	return ans
}

/*
从三个不发送好友请求的条件来看，以 y 的角度来说，可总结为：年龄比我小的不考虑（同龄的可以），年龄比我大可以考虑，但是不能超过一定范围则不考虑。
即对于一个确定的 y 而言，会发送好友请求的 x 范围为连续段：
*/

func numFriendRequests(ages []int) int {
	ans := 0
	sort.Ints(ages)
	le, ri := 0, 0
	for k := 0; k < len(ages); k++ {
		// 找左边

		for le-1 >= 0 && !(ages[le-1] <= ages[k]/2+7) {
			le--
		}
		// 找右边
		for ri+1 < len(ages) && ages[ri+1] == ages[k] {
			ri++
		}
		// 边界处理太难了
		// if ri > le {
		ans += ri - k + (k - le)
		// }
	}

	return ans
}

func numFriendRequests4(ages []int) (ans int) {
	sort.Ints(ages)
	left, right := 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		ans += right - left
	}
	return
}

// 检测 x 是否向 y 发信息
func check2(x, y int) bool {
	if x == y {
		return false
	}
	if y <= x/2+7 {
		return false
	}
	if y > x {
		return false
	}

	if y > 100 && x < 100 {
		return false
	}

	return true
}

func check(x, y int) bool {
	if y <= x/2+7 {
		return false
	}
	if y > x {
		return false
	}
	if y > 100 && x < 100 {
		return false
	}
	return true
}

/*
如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：
ages[y] <= 0.5 * ages[x] + 7
ages[y] > ages[x]
ages[y] > 100 && ages[x] < 100
*/
