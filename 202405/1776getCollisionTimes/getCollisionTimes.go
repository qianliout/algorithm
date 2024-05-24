package main

import (
	"fmt"
)

func main() {
	fmt.Println(getCollisionTimes([][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}))
}

func getCollisionTimes(cars [][]int) []float64 {
	// 满足从栈底到栈顶的下标对应的相遇时间单调递减
	st := make([]int, 0)

	ans := make([]float64, len(cars))

	for i := len(cars) - 1; i >= 0; i-- {
		if len(st) == 0 {
			ans[i] = -1
			st = append(st, i)
			continue
		}

		/*
			如果栈不为空且当前车不可能与栈顶车相遇，则将栈顶下标出栈。当前车不可能与栈顶车相遇可能有如下情况。
			1 栈顶车的速度大于等于当前车的速度。
			2 栈顶车与更前面的车相遇且相遇时间小于等于当前车与栈顶车相遇的时间，则在当前车与栈顶车相遇的同时或更早的时候栈顶车
			与更前面的车属于同一个车队，因此当前车将与更前面的车相遇。
		*/

		for len(st) > 0 && (cars[i][1] <= cars[st[len(st)-1]][1] || (ans[st[len(st)-1]] > 0 && ans[st[len(st)-1]] <= getCollision(cars[i], cars[st[len(st)-1]]))) {
			st = st[:len(st)-1]
		}

		/*
			如果栈不为空，则当前车将与栈顶车相遇，计算当前车与栈顶车相遇的时间并填入 ans[i]；如果栈为空，则当前车不会与任何车相遇，将 −1 填入 ans[i]
		*/
		if len(st) > 0 {
			ans[i] = getCollision(cars[i], cars[st[len(st)-1]])
		} else {
			ans[i] = -1
		}
		st = append(st, i)
	}
	return ans
}

// 大于0,能追上
func getCollision(currCar []int, prevCar []int) float64 {
	return float64(prevCar[0]-currCar[0]) / float64(currCar[1]-prevCar[1])
}
