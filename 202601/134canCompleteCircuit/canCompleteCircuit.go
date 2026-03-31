package main

import (
	"fmt"
)

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 2, 1}
	ans := canCompleteCircuit(gas, cost)
	fmt.Println("ans is ", ans)
}

func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	start := 0
	all := 0
	n := len(gas)
	for i := 0; i < n; i++ {
		if all < 0 {
			start = i
			all = 0
		}
		all = all + gas[i] - cost[i]
	}

	return start % n
}

func sum(num []int) int {
	ans := 0
	for _, n := range num {
		ans += n
	}
	return ans
}

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/
